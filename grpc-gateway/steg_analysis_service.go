package main

import (
	"context"
	"fmt"
	pb "grpcgw/pb"
	"hash/adler32"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Knetic/govaluate"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v2"
)

type IStegAnalysisService interface {
	Execute(context.Context, *pb.StegAnalysisRequest) (*pb.StegAnalysisResponse, error)
	LoadWorkflowConfig() error
	GetTaskByTaskId(string) (Task, error)
	HideOutput() bool
}

type Task struct {
	FuncName        string
	Semaphore       chan struct{}
	TaskId          string                 `yaml:"name"`
	SvcName         string                 `yaml:"exec"`
	File            string                 `yaml:"file"`
	Cond            string                 `yaml:"cond"`
	HiddenFields    []string               `yaml:"hide"`
	VisibleFields   []string               `yaml:"return"`
	AsyncTasks      []Task                 `yaml:"async"`
	SyncTasks       []Task                 `yaml:"sync"`
	IteratorTasks   []Task                 `yaml:"iter"`
	Range           string                 `yaml:"range"`
	HideResultOnErr bool                   `yaml:"hide_on_err"`
	Params          map[string]interface{} `yaml:"param"`
	TimeoutInSec    float32                `yaml:"task_timeout_in_sec"`
	MaxFileSzInKb   int                    `yaml:"max_file_size_in_kb"`
	MaxConcurReq    int                    `yaml:"max_concurrent_req"`
	AllwdFlTypes    []string               `yaml:"allowed_file_types"`
	AdditionalVals  map[string]interface{} `yaml:"additional_return_fields"`
	ShowOutput      bool                   `yaml:"show_output"`
}

type StegAnalysisService struct {
	ctxGrpcSvr       *GrpcServerContext
	semaphore        chan struct{}
	goValuateFuncs   map[string]govaluate.ExpressionFunction
	varsInUse        []string
	SyncTasks        []Task   `yaml:"sync"`
	AsyncTasks       []Task   `yaml:"async"`
	WfTimeoutInSec   float32  `yaml:"workflow_timeout_in_sec"`
	TaskTimeoutInSec float32  `yaml:"task_timeout_in_sec"`
	MaxFileSzInKb    int      `yaml:"max_file_size_in_kb"`
	MaxConcurReq     int      `yaml:"max_concurrent_req"`
	AllwdFlTypes     []string `yaml:"allowed_file_types"`
	ToggleOutput     bool     `yaml:"toggle_output"`
}

func NewStegAnalysisService(ctx *GrpcServerContext) IStegAnalysisService {
	stdWfSvc := StegAnalysisService{ctxGrpcSvr: ctx}
	if err := stdWfSvc.LoadWorkflowConfig(); err != nil {
		log.Printf("WARNING: could not load workflow: %v", err)
	}
	return &stdWfSvc
}

func setTaskProperties(tasks *[]Task, timeoutInSec float32, tskCnt int, iterCnt int, asyncCnt int, syncCnt int, showOutput bool, toggleOutput bool) error {
	for i, v := range *tasks {
		if v.AsyncTasks != nil {
			asyncCnt++
			(*tasks)[i].TaskId = fmt.Sprintf("a%v", asyncCnt)
			err := setTaskProperties(&v.AsyncTasks, timeoutInSec, tskCnt, iterCnt, asyncCnt, syncCnt, v.ShowOutput == toggleOutput, toggleOutput)
			if err != nil {
				return err
			}
			continue
		} else if v.SyncTasks != nil {
			syncCnt++
			(*tasks)[i].TaskId = fmt.Sprintf("s%v", syncCnt)
			err := setTaskProperties(&v.SyncTasks, timeoutInSec, tskCnt, iterCnt, asyncCnt, syncCnt, v.ShowOutput == toggleOutput, toggleOutput)
			if err != nil {
				return err
			}
			continue
		} else if v.IteratorTasks != nil {
			iterCnt++
			(*tasks)[i].TaskId = fmt.Sprintf("i%v", iterCnt)
			err := setTaskProperties(&v.IteratorTasks, timeoutInSec, tskCnt, iterCnt, asyncCnt, syncCnt, v.ShowOutput == toggleOutput, toggleOutput)
			if err != nil {
				return err
			}
			continue
		} else {
			tskCnt++

			// set task id if name was not set
			if (*tasks)[i].TaskId == "" {
				(*tasks)[i].TaskId = fmt.Sprintf("t%v", tskCnt)
			}

			// set func name
			if strings.Count(v.SvcName, ".") != 1 {
				return fmt.Errorf("%v is no valid function call", v.SvcName)
			}
			(*tasks)[i].SvcName = strings.Split(v.SvcName, ".")[0]
			(*tasks)[i].FuncName = strings.Split(v.SvcName, ".")[1]

			// overwrite max_concurrent_req if set in task
			if (*tasks)[i].MaxConcurReq > 0 {
				(*tasks)[i].Semaphore = make(chan struct{}, (*tasks)[i].MaxConcurReq)
			}

			// set timeout
			if timeoutInSec > 0 && (*tasks)[i].TimeoutInSec == 0 {
				(*tasks)[i].TimeoutInSec = timeoutInSec
			}

			// set hide task result
			if !(*tasks)[i].ShowOutput {
				(*tasks)[i].ShowOutput = showOutput
			}
		}
	}
	return nil
}

func addVarFromCond(svc *StegAnalysisService, tasks []Task) {
	for _, t := range tasks {
		if t.AsyncTasks != nil {
			addVarFromCond(svc, t.AsyncTasks)
		}
		if t.SyncTasks != nil {
			addVarFromCond(svc, t.AsyncTasks)
		}
		if t.Cond != "" {
			re := regexp.MustCompile(`([\w.]+->[\w.]+(?:\[\d+\])?(?:\.[\w.]+(?:\[\d+\])?)*)`)
			matches := re.FindAllString(t.Cond, -1)
			svc.varsInUse = append(svc.varsInUse, matches...)
		}
	}
}

func processValues(taskId string, response *pb.StegServiceResponse, vars *sync.Map) {
	// iterate over response values recursive
	for key, value := range response.Values {
		addValueToVars(fmt.Sprintf("%s->%s", taskId, key), value, vars)
	}
}

func addValueToVars(valueName string, value *pb.ResponseValue, vars *sync.Map) {
	if value == nil {
		vars.Store(valueName, nil)
		return
	}

	switch t := value.Value.(type) {
	case *pb.ResponseValue_StringValue:
		vars.Store(valueName, t.StringValue)
	case *pb.ResponseValue_IntValue:
		vars.Store(valueName, t.IntValue)
	case *pb.ResponseValue_FloatValue:
		vars.Store(valueName, float64(t.FloatValue))
	case *pb.ResponseValue_BoolValue:
		vars.Store(valueName, t.BoolValue)
	case *pb.ResponseValue_BinaryValue:
		if len(t.BinaryValue) == 0 {
			vars.Store(valueName, nil)
		} else {
			vars.Store(valueName, t.BinaryValue)
		}
	case *pb.ResponseValue_StructuredValue:
		switch kind := t.StructuredValue.Kind.(type) {
		case *structpb.Value_StringValue:
			vars.Store(valueName, kind.StringValue)
		case *structpb.Value_NumberValue:
			vars.Store(valueName, kind.NumberValue)
		case *structpb.Value_BoolValue:
			vars.Store(valueName, kind.BoolValue)
		case *structpb.Value_ListValue:
			for i, listItem := range kind.ListValue.Values {
				itemName := fmt.Sprintf("%s[%d]", valueName, i)
				addValueToVars(itemName, &pb.ResponseValue{
					Value: &pb.ResponseValue_StructuredValue{StructuredValue: listItem},
				}, vars)
			}

			if len(kind.ListValue.Values) == 0 {
				vars.Store(valueName, nil)
			} else {
				vars.Store(valueName, t)
			}
		case *structpb.Value_StructValue:
			for k, v := range kind.StructValue.Fields {
				addValueToVars(fmt.Sprintf("%s.%s", valueName, k), &pb.ResponseValue{
					Value: &pb.ResponseValue_StructuredValue{StructuredValue: v},
				}, vars)
			}

			if len(kind.StructValue.Fields) == 0 {
				vars.Store(valueName, nil)
			} else {
				vars.Store(valueName, t)
			}
		case *structpb.Value_NullValue:
			vars.Store(valueName, nil)
		}
	default:
		fmt.Printf("Unsupported type for value: %T\n", t)
	}
}

func getAllowedFileTypesOfSvcFunc(svcName string, funcName string, svcs []StegService) []string {
	for _, v := range svcs {
		if v.svcInfo.Name == svcName {
			for _, f := range v.svcInfo.Functions {
				if f.Name == funcName {
					return f.SupportedFileTypes
				}
			}
		}
	}
	return nil
}

func concatParams(clientParams map[string]string, wfParams map[string]interface{}, vars *sync.Map, svcNameOpt ...string) (map[string]*pb.StegServiceRequestParameterValue, error) {
	reqParams := make(map[string]*pb.StegServiceRequestParameterValue)

	for key, value := range wfParams {
		if strings.Count(key, "->") == 1 {
			paramName := strings.Split(key, "->")[1]

			val, ok := vars.Load(key)
			if !ok {
				return nil, fmt.Errorf("variable %s not existing", key)
			}
			switch v := val.(type) {
			case string:
				reqParams[paramName] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_StringValue{StringValue: v}}
			case int:
				reqParams[paramName] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_IntValue{IntValue: int64(v)}}
			case float32:
				reqParams[paramName] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_FloatValue{FloatValue: v}}
			case bool:
				reqParams[paramName] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_BoolValue{BoolValue: v}}
			case []byte:
				reqParams[paramName] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_BinaryValue{BinaryValue: v}}
			default:
				return nil, fmt.Errorf("unsupported type for variable %q: %T", key, value)
			}
		} else {
			switch v := value.(type) {
			case string:
				reqParams[key] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_StringValue{StringValue: v}}
			case int:
				reqParams[key] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_IntValue{IntValue: int64(v)}}
			case float32:
				reqParams[key] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_FloatValue{FloatValue: v}}
			case bool:
				reqParams[key] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_BoolValue{BoolValue: v}}
			case []byte:
				reqParams[key] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_BinaryValue{BinaryValue: v}}
			default:
				return nil, fmt.Errorf("unsupported type for variable %q: %T", key, value)
			}
		}
	}

	for key, value := range clientParams {
		if len(svcNameOpt) > 0 {
			fmtSvcName := fmt.Sprintf("%v->", svcNameOpt[0])
			if strings.Contains(key, fmtSvcName) {
				reqParams[strings.Replace(key, fmtSvcName, "", 1)] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_StringValue{StringValue: value}}
				continue
			}
		}
		reqParams[key] = &pb.StegServiceRequestParameterValue{Value: &pb.StegServiceRequestParameterValue_StringValue{StringValue: value}}
	}

	return reqParams, nil
}

func convertReturnValueToBytes(paramName string, vars *sync.Map) ([]byte, error) {
	if strings.Count(paramName, "->") == 1 {
		val, ok := vars.Load(paramName)
		if !ok {
			return nil, fmt.Errorf("no variable %v existing", paramName)
		}
		switch v := val.(type) {
		case []byte:
			return v, nil
		default:
			return nil, fmt.Errorf("variable %v not type bytes", paramName)
		}
	}
	return nil, fmt.Errorf("parameter %v not a valid variable", paramName)
}

func setEmptyVars(tasks []Task, vars *sync.Map, svcs []StegService, taskResponses *[]*pb.TaskResult) {
	for _, t := range tasks {
		if t.AsyncTasks != nil {
			setEmptyVars(t.AsyncTasks, vars, svcs, taskResponses)
		}
		if t.SyncTasks != nil {
			setEmptyVars(t.SyncTasks, vars, svcs, taskResponses)
		}
		if t.IteratorTasks != nil {
			setEmptyVars(t.IteratorTasks, vars, svcs, taskResponses)
		}
		if t.SvcName != "" {
			// get related service
			s, errs := GetServiceByName(svcs, t.SvcName)
			if errs != nil {
				continue
			}

			f, errs := GetFunctionByName(s, t.FuncName)
			if errs != nil {
				continue
			}

			for _, retFields := range f.ReturnFields {
				_, ok := vars.Load(fmt.Sprintf("%s->%s", t.TaskId, retFields.Name))
				if !ok {
					vars.Store(fmt.Sprintf("%s->%s", t.TaskId, retFields.Name), nil)
				}
			}
		}
		if t.AdditionalVals != nil {
			for i := range t.AdditionalVals {
				_, ok := vars.Load(fmt.Sprintf("%s->%s", t.TaskId, i))
				if !ok {
					vars.Store(fmt.Sprintf("%s->%s", t.TaskId, i), nil)
				}
			}
		}
	}
}

func setResponse(status pb.Status, svcResp *pb.StegServiceResponse, err string, duration int64, task Task, taskResponses *[]*pb.TaskResult, vars *sync.Map, svcs []StegService) {
	vars.Store(fmt.Sprintf("%s->error", task.TaskId), err)

	// enrich empty return values
	if task.SvcName != "" {
		// get related service
		s, errs := GetServiceByName(svcs, task.SvcName)
		if errs != nil {
			log.Print(errs.Error())
		} else {
			f, errs := GetFunctionByName(s, task.FuncName)
			if errs != nil {
				log.Print(errs.Error())
			} else {
				for _, retFields := range f.ReturnFields {
					_, ok := vars.Load(fmt.Sprintf("%s->%s", task.TaskId, retFields.Name))
					if !ok {
						vars.Store(fmt.Sprintf("%s->%s", task.TaskId, retFields.Name), nil)
					}
				}
			}
		}
	} else {
		if task.AsyncTasks != nil {
			setEmptyVars(task.AsyncTasks, vars, svcs, taskResponses)
		}
		if task.SyncTasks != nil {
			setEmptyVars(task.SyncTasks, vars, svcs, taskResponses)
		}
		if task.IteratorTasks != nil {
			setEmptyVars(task.IteratorTasks, vars, svcs, taskResponses)
		}
	}

	*taskResponses = append(*taskResponses, &pb.TaskResult{TaskId: task.TaskId, ServiceName: task.SvcName, FunctionName: task.FuncName, Error: err, ServiceResponse: svcResp, Status: status.String(), DurationMs: duration})
}

func getServiceFunction(svcName string, funcName string, svcs []StegService) (*pb.StegServiceFunction, error) {
	for _, svc := range svcs {
		if svc.svcInfo.Name == svcName {
			for _, fn := range svc.svcInfo.Functions {
				if fn.Name == funcName {
					return fn, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("service %s doesnt contain function %s", svcName, funcName)
}

func parseRange(rangeStr string, vars *sync.Map, funcs map[string]govaluate.ExpressionFunction) (interface{}, error) {
	regex := `^(\d+(\.\d+)?) (\d+(\.\d+)?)( (\d+(\.\d+)?))?$` // number range
	r := regexp.MustCompile(regex)

	matches := r.FindStringSubmatch(rangeStr)
	if matches != nil {
		start, _ := strconv.ParseFloat(matches[1], 64)
		end, _ := strconv.ParseFloat(matches[3], 64)
		step, err := strconv.ParseFloat(matches[6], 64)
		if err != nil {
			step = 1
		}
		var iterVals []interface{}
		for value := start; value < end; value += step {
			iterVals = append(iterVals, value)
		}
		return iterVals, nil
	} else {
		result, err := parseGoEvalExpr(rangeStr, syncMapToInterfaceMap(vars), funcs)
		if err != nil {
			return nil, fmt.Errorf("error parsing range expression: %v", err)
		}
		switch t := result.(type) {
		case []string, []int, []float64, []float32, []bool, []*structpb.Value, map[string]string, map[string]int, map[string]float64, map[string]float32, map[string]bool, map[string]*structpb.Value:
			return result, nil
		case *pb.ResponseValue_StructuredValue:
			switch kind := t.StructuredValue.Kind.(type) {
			case *structpb.Value_StringValue, *structpb.Value_NumberValue, *structpb.Value_BoolValue, *structpb.Value_NullValue:
				return nil, fmt.Errorf("type %v of variable %v not iterable", t, result)
			case *structpb.Value_ListValue:
				return kind.ListValue.Values, nil
			case *structpb.Value_StructValue:
				return kind.StructValue.Fields, nil
			}
			return result, nil
		default:
			return nil, fmt.Errorf("type %v of variable %v not iterable", t, result)
		}
	}
}

func removeKeys(keysToRm []string, vars *sync.Map) {
	for _, keyBase := range keysToRm {
		vars.Range(func(key, value interface{}) bool {
			keyStr, ok := key.(string)
			if !ok {
				return true
			}
			if keyStr == keyBase || strings.HasPrefix(keyStr, fmt.Sprintf("%v->", keyBase)) {
				vars.Delete(keyStr)
			}
			return true
		})
	}
}

func getTaskByTaskId(tasks []Task, taskId string) (Task, error) {
	for _, t := range tasks {
		if t.AsyncTasks != nil {
			task, err := getTaskByTaskId(t.AsyncTasks, taskId)
			if err == nil {
				return task, nil
			}
		}
		if t.SyncTasks != nil {
			task, err := getTaskByTaskId(t.SyncTasks, taskId)
			if err == nil {
				return task, nil
			}
		}
		if t.IteratorTasks != nil {
			task, err := getTaskByTaskId(t.IteratorTasks, taskId)
			if err == nil {
				return task, nil
			}
		}
		if t.TaskId == taskId {
			return t, nil
		}
	}
	return Task{}, fmt.Errorf("no task %v found", taskId)
}

func syncMapToInterfaceMap(sMap *sync.Map) map[string]interface{} {
	result := make(map[string]interface{})
	sMap.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})
	return result
}

func (svc *StegAnalysisService) Execute(ctx context.Context, req *pb.StegAnalysisRequest) (*pb.StegAnalysisResponse, error) {
	// wait at semaphore if enabled
	if svc.semaphore != nil {
		maxSemTimeout := 30 * time.Second
		ctx, cancel := context.WithTimeout(ctx, maxSemTimeout)
		defer cancel()

		select {
		case svc.semaphore <- struct{}{}:
			defer func() {
				<-svc.semaphore
			}()
		case <-ctx.Done():
			return nil, fmt.Errorf("semaphore timeout: %v", ctx.Err())
		}
	}

	// check file size
	if len(req.File)/1024 > svc.MaxFileSzInKb && svc.MaxFileSzInKb != 0 {
		return nil, fmt.Errorf("file (%v KB) exeeds limit of %v", len(req.File)/1024, svc.MaxFileSzInKb)
	}

	// check file type
	if svc.AllwdFlTypes != nil {
		if allwd, flType := checkFileTypeAllowed(req.File, svc.AllwdFlTypes); !allwd {
			return nil, fmt.Errorf("file type %v not allowed", flType)
		}
	}

	// set workflow timeout
	if svc.WfTimeoutInSec > 0 {
		wfTimeout := time.Duration(svc.WfTimeoutInSec*1000) * time.Millisecond
		wfctx, cancel := context.WithTimeout(ctx, wfTimeout)
		defer cancel()
		ctx = wfctx
	}

	taskResps := make([]*pb.TaskResult, 0)
	wfVars := sync.Map{}

	if req.Exec != "" {
		if strings.Count(req.Exec, ".") != 1 {
			return nil, fmt.Errorf("%v is no valid function call", req.Exec)
		}
		svcName := strings.Split(req.Exec, ".")[0]
		funcName := strings.Split(req.Exec, ".")[1]
		t := Task{TaskId: req.Exec, SvcName: svcName, FuncName: funcName}
		svc.executeStep(ctx, t, svc.ctxGrpcSvr.stegSvcs, &taskResps, req, &wfVars)
	} else if svc.SyncTasks != nil {
		for _, task := range svc.SyncTasks {
			// set standard timeout if not set
			if task.TimeoutInSec == 0 {
				task.TimeoutInSec = svc.TaskTimeoutInSec
			}
			svc.executeStep(ctx, task, svc.ctxGrpcSvr.stegSvcs, &taskResps, req, &wfVars)
		}
	} else if svc.AsyncTasks != nil {
		for _, task := range svc.AsyncTasks {
			// set standard timeout if not set
			if task.TimeoutInSec == 0 {
				task.TimeoutInSec = svc.TaskTimeoutInSec
			}
			svc.executeStep(ctx, task, svc.ctxGrpcSvr.stegSvcs, &taskResps, req, &wfVars)
		}
	} else {
		return nil, fmt.Errorf("workflow missing exec group")
	}

	saResp := pb.StegAnalysisResponse{TaskResults: taskResps}
	return &saResp, nil
}

func (svc *StegAnalysisService) executeStep(ctx context.Context, task Task, svcs []StegService, taskResults *[]*pb.TaskResult, req *pb.StegAnalysisRequest, vars *sync.Map) {
	// handle conditions
	if task.Cond != "" {
		// evaluate expression
		result, err := parseGoEvalExpr(task.Cond, syncMapToInterfaceMap(vars), svc.goValuateFuncs)
		if err != nil {
			setResponse(pb.Status_OUT_OF_CONDITION, nil, err.Error(), 0, task, taskResults, vars, svcs)
			return
		}

		if !result.(bool) {
			setResponse(pb.Status_OUT_OF_CONDITION, nil, "", 0, task, taskResults, vars, svcs)
			return
		}
	}

	// semaphore for max requests per task
	if task.Semaphore != nil {
		select {
		case task.Semaphore <- struct{}{}:
			defer func() { <-task.Semaphore }()
		case <-time.After(30 * time.Second):
			setResponse(pb.Status_EXCEEDED_TIMEOUT, nil, "semaphore timeout", 0, task, taskResults, vars, svcs)
			return
		case <-ctx.Done():
			setResponse(pb.Status_GRPC_ERROR, nil, ctx.Err().Error(), 0, task, taskResults, vars, svcs)
			return
		}
	}

	// check max file size
	if len(req.File)/1024 > task.MaxFileSzInKb && task.MaxFileSzInKb != 0 {
		setResponse(pb.Status_EXCEEDED_FILESIZE, nil, "", 0, task, taskResults, vars, svcs)
		return
	}

	// check file type
	allwdFileTypes := getAllowedFileTypesOfSvcFunc(task.SvcName, task.FuncName, svcs)
	if allwdFileTypes != nil {
		if allwd, _ := checkFileTypeAllowed(req.File, allwdFileTypes); !allwd {
			setResponse(pb.Status_FILE_TYPE_NOT_ALLOWED, nil, "", 0, task, taskResults, vars, svcs)
			return
		}
	}
	if task.AllwdFlTypes != nil {
		if allwd, _ := checkFileTypeAllowed(req.File, task.AllwdFlTypes); !allwd {
			setResponse(pb.Status_FILE_TYPE_NOT_ALLOWED, nil, "", 0, task, taskResults, vars, svcs)
			return
		}
	}

	// async tasks
	if len(task.AsyncTasks) > 0 {
		var wg sync.WaitGroup
		for _, asyncStep := range task.AsyncTasks {
			wg.Add(1)
			go func() {
				defer wg.Done()
				svc.executeStep(ctx, asyncStep, svcs, taskResults, req, vars)
			}()
		}
		wg.Wait()
		return
	}

	// sync tasks
	if len(task.SyncTasks) > 0 {
		for _, syncStep := range task.SyncTasks {
			svc.executeStep(ctx, syncStep, svcs, taskResults, req, vars)
		}
		return
	}

	// iterator tasks
	if len(task.IteratorTasks) > 0 {
		values, err := parseRange(task.Range, vars, svc.goValuateFuncs)
		if err != nil {
			setResponse(pb.Status_GRPC_ERROR, nil, fmt.Sprintf("error parsing range: %v", err), 0, task, taskResults, vars, svcs)
			return
		}

		switch tValues := values.(type) {
		case []interface{}:
			for _, v := range tValues {
				vars.Store(task.TaskId, v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case []string:
			for _, v := range tValues {
				vars.Store(task.TaskId, v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case []int:
			for _, v := range tValues {
				vars.Store(task.TaskId, v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case []float64:
			for _, v := range tValues {
				vars.Store(task.TaskId, v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case []float32:
			for v := range tValues {
				vars.Store(task.TaskId, v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case []bool:
			for v := range tValues {
				vars.Store(task.TaskId, v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case []*structpb.Value:
			for _, v := range tValues {
				vars.Store(task.TaskId, v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case map[string]interface{}:
			for k, v := range tValues {
				vars.Store(fmt.Sprintf("%v->key", task.TaskId), k)
				vars.Store(fmt.Sprintf("%v->value", task.TaskId), v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case map[string]string:
			for k, v := range tValues {
				vars.Store(fmt.Sprintf("%v->key", task.TaskId), k)
				vars.Store(fmt.Sprintf("%v->value", task.TaskId), v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case map[string]int:
			for k, v := range tValues {
				vars.Store(fmt.Sprintf("%v->key", task.TaskId), k)
				vars.Store(fmt.Sprintf("%v->value", task.TaskId), v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case map[string]float64:
			for k, v := range tValues {
				vars.Store(fmt.Sprintf("%v->key", task.TaskId), k)
				vars.Store(fmt.Sprintf("%v->value", task.TaskId), v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case map[string]float32:
			for k, v := range tValues {
				vars.Store(fmt.Sprintf("%v->key", task.TaskId), k)
				vars.Store(fmt.Sprintf("%v->value", task.TaskId), v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case map[string]bool:
			for k, v := range tValues {
				vars.Store(fmt.Sprintf("%v->key", task.TaskId), k)
				vars.Store(fmt.Sprintf("%v->value", task.TaskId), v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		case map[string]*structpb.Value:
			for k, v := range tValues {
				vars.Store(fmt.Sprintf("%v->key", task.TaskId), k)
				vars.Store(fmt.Sprintf("%v->value", task.TaskId), v)
				tmpTskIds := make([]string, 0)
				for _, iterStep := range task.IteratorTasks {
					tmpTskIds = append(tmpTskIds, iterStep.TaskId)
					svc.executeStep(ctx, iterStep, svcs, taskResults, req, vars)
				}
				removeKeys(tmpTskIds, vars)
			}
		default:
			setResponse(pb.Status_GRPC_ERROR, nil, fmt.Sprintf("range type %v not supported", tValues), 0, task, taskResults, vars, svcs)
			return
		}
		return
	}

	// no service name specified
	if task.SvcName == "" {
		setResponse(pb.Status_CLIENT_ERROR, nil, "no service name specified", 0, task, taskResults, vars, svcs)
		return
	}

	// get related service
	s, err := GetServiceByName(svcs, task.SvcName)
	if err != nil {
		setResponse(pb.Status_GRPC_ERROR, nil, err.Error(), 0, task, taskResults, vars, svcs)
		return
	}

	// set timeout if specified
	if task.TimeoutInSec > 0 {
		taskTimeout := time.Duration(task.TimeoutInSec*1000) * time.Millisecond
		wfctx, cancel := context.WithTimeout(ctx, taskTimeout)
		defer cancel()
		ctx = wfctx
	}

	// enrich request params with task params
	newParams, err := concatParams(req.Params, task.Params, vars, task.SvcName)
	if err != nil {
		setResponse(pb.Status_GRPC_ERROR, nil, err.Error(), 0, task, taskResults, vars, svcs)
		return
	}

	fn, err := getServiceFunction(task.SvcName, task.FuncName, svcs)
	if err != nil {
		setResponse(pb.Status_GRPC_ERROR, nil, err.Error(), 0, task, taskResults, vars, svcs)
		return
	}

	// check if File property was overwritten
	var file []byte
	if !fn.FileOptional {
		file = req.File
		if task.File != "" {
			file, err = convertReturnValueToBytes(task.File, vars)
			if err != nil {
				setResponse(pb.Status_GRPC_ERROR, nil, err.Error(), 0, task, taskResults, vars, svcs)
				return
			}
		}
	}

	// execute service
	start := time.Now()
	var res *pb.StegServiceResponse

	if fn.IsNop {
		res = &pb.StegServiceResponse{}
		err = nil
	} else {
		res, err = s.Execute(ctx, &pb.StegServiceRequest{File: file, Function: task.FuncName, Params: newParams})
	}

	duration := time.Since(start).Milliseconds()
	if err != nil {
		setResponse(pb.Status_SERVICE_ERROR, nil, err.Error(), duration, task, taskResults, vars, svcs)
		return
	}
	if res.Error != "" {
		setResponse(pb.Status_SERVICE_ERROR, nil, res.Error, duration, task, taskResults, vars, svcs)
		return
	}

	if res.Values == nil {
		res.Values = make(map[string]*pb.ResponseValue)
	}

	// add additional values (print statement)
	if task.AdditionalVals != nil {
		for k, v := range task.AdditionalVals {
			// evaluate expression
			result, err := parseGoEvalExpr(fmt.Sprintf("%v", v), syncMapToInterfaceMap(vars), svc.goValuateFuncs)
			if err != nil {
				log.Printf("WARNING: could not evaluate expression, passing as constant: %v", v)
				result = v
			}
			switch r := result.(type) {
			case string:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_StringValue{StringValue: r}}
			case int:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_IntValue{IntValue: int64(r)}}
			case float64:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_FloatValue{FloatValue: float32(r)}}
			case float32:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_FloatValue{FloatValue: r}}
			case bool:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_BoolValue{BoolValue: r}}
			case []byte:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_BinaryValue{BinaryValue: r}}
			case *pb.ResponseValue_StructuredValue:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_StructuredValue{StructuredValue: r.StructuredValue}}
			case *structpb.Value:
				res.Values[k] = &pb.ResponseValue{Value: &pb.ResponseValue_StructuredValue{StructuredValue: r}}
			default:
				setResponse(pb.Status_SERVICE_ERROR, nil, fmt.Sprintf("Error evaluating statement for %v: type %T not supported", k, r), duration, task, taskResults, vars, svcs)
				return
			}
		}
	}

	// append return values to vars
	processValues(task.TaskId, res, vars)

	// apply filter for return/hidden fields
	if task.VisibleFields != nil {
		// include visible fields (return fields)
		for key := range res.Values {
			if !slices.Contains(task.VisibleFields, key) {
				delete(res.Values, key)
			}
		}
	} else if task.HiddenFields != nil {
		// exclude hidden fields
		for key := range res.Values {
			if slices.Contains(task.HiddenFields, key) {
				delete(res.Values, key)
			}
		}
	}

	// append current service result
	setResponse(pb.Status_SUCCESS, res, "", duration, task, taskResults, vars, svcs)
}
func (svc *StegAnalysisService) HideOutput() bool {
	return svc.ToggleOutput
}

func (svc *StegAnalysisService) LoadWorkflowConfig() error {
	*svc = StegAnalysisService{
		ctxGrpcSvr: svc.ctxGrpcSvr,
	}

	data, err := os.ReadFile(WfFileName)
	if err != nil {
		return fmt.Errorf("error reading %v: %v", WfFileName, err)
	}

	// calc hash
	hasher := adler32.New()
	hasher.Write(data)
	WfFileHash = hasher.Sum32()

	if err = yaml.Unmarshal(data, &svc); err != nil {
		return fmt.Errorf("error parsing %v: %v", WfFileName, err)
	}

	if (svc.AsyncTasks != nil && svc.SyncTasks != nil) ||
		(svc.AsyncTasks == nil && svc.SyncTasks == nil) {
		return fmt.Errorf("please specify a valid workflow")
	}

	if svc.AsyncTasks != nil {
		if err := setTaskProperties(&svc.AsyncTasks, svc.TaskTimeoutInSec, 0, 0, 0, 0, svc.ToggleOutput, svc.ToggleOutput); err != nil {
			return err
		}
		addVarFromCond(svc, svc.AsyncTasks)
	} else {
		if err := setTaskProperties(&svc.SyncTasks, svc.TaskTimeoutInSec, 0, 0, 0, 0, svc.ToggleOutput, svc.ToggleOutput); err != nil {
			return err
		}
		addVarFromCond(svc, svc.SyncTasks)
	}

	if svc.MaxConcurReq > 0 {
		svc.semaphore = make(chan struct{}, svc.MaxConcurReq)
	}

	svc.goValuateFuncs = getGoValuateFuncs()

	OutputDir, err = createOutputFolder()
	if err != nil {
		return fmt.Errorf("error creating output dir: %v", err)
	}

	return nil
}

func (svc *StegAnalysisService) GetTaskByTaskId(taskId string) (Task, error) {
	if svc.AsyncTasks != nil {
		return getTaskByTaskId(svc.AsyncTasks, taskId)
	}
	if svc.SyncTasks != nil {
		return getTaskByTaskId(svc.SyncTasks, taskId)
	}
	return Task{}, fmt.Errorf("no task found for %v", taskId)
}
