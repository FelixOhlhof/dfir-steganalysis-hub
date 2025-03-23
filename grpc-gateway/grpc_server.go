package main

import (
	"context"
	"fmt"
	pb "grpcgw/pb"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var OutputDir string

type GrpcServerContext struct {
	gwAddr   string
	stegSvcs []StegService
}

type GrpcServer struct {
	ctx      GrpcServerContext
	stegaSvc IStegAnalysisService
	confSvc  IConfigService
	rptSvc   RptSvc
	reqCnt   int
	pb.UnimplementedStegAnalysisServiceServer
	pb.UnimplementedConfigServiceServer
}

func NewGrpcServer() *GrpcServer {
	ctx := GrpcServerContext{}
	rptSvc := NewRptSvc()
	rptSvc.Start()

	// create config service
	confsvc := NewConfigService(&ctx)
	confsvc = NewConfLoggingService(confsvc)

	// create new workflow service
	stegSvc := NewStegAnalysisService(&ctx)
	// stegSvc = NewStegLoggingService(stegSvc)
	// stegSvc = NewStegReportService(stegSvc)

	// initialize context
	ctx.initialize()

	return &GrpcServer{ctx: ctx, stegaSvc: stegSvc, confSvc: confsvc, rptSvc: rptSvc, reqCnt: 0}
}

func (ctx *GrpcServerContext) initialize() {
	// set gw address
	gwPort := os.Getenv("port")
	if gwPort == "" {
		gwPort = "50000"
	}
	if !strings.Contains(gwPort, ":") {
		gwPort = fmt.Sprintf(":%v", gwPort)
	}
	ctx.gwAddr = gwPort

	// connect to registered services
	stegSvcs := strings.TrimRight(os.Getenv("services"), ";")
	tmp := strings.Split(stegSvcs, ";")
	for _, addr := range tmp {
		svc := NewStegServiceClient(addr)
		if err := svc.Connect(); err != nil {
			log.Printf("WARNING: could not connect to %v: %v", addr, err)
		} else {
			ctx.stegSvcs = append(ctx.stegSvcs, svc)
		}
	}

	for _, v := range ctx.stegSvcs {
		log.Printf("connected to %v at %v", v.svcInfo.Name, v.addr)
	}
}

func createOutputFolder() (string, error) {
	os.MkdirAll(OutputDirName, os.ModePerm)
	entries, err := os.ReadDir(OutputDirName)
	if err != nil {
		return "", fmt.Errorf("error reading while reading report directory %v", err)
	}

	lastCount := 0
	lastCheckSum := ""
	lastDir := ""

	for _, entry := range entries {
		if entry.IsDir() && strings.Contains(entry.Name(), "_") {
			cnt, _ := strconv.Atoi(strings.Split(entry.Name(), "_")[0])
			cKsum := strings.Split(entry.Name(), "_")[1]
			if cnt > lastCount {
				lastCount = cnt
				lastCheckSum = cKsum
				lastDir = filepath.Join(OutputDirName, entry.Name())
			}
		}
	}

	curHash := fmt.Sprintf("%v", WfFileHash)
	if lastCheckSum == curHash {
		return lastDir, nil
	}

	folderName := fmt.Sprintf("%v_%v", lastCount+1, curHash)
	path := filepath.Join(OutputDirName, folderName)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}

	if err := copyFileToFolder(WfFileName, path); err != nil {
		return "", err
	}

	return path, nil
}

func convertTaskResultPointersToValues(taskResults []*pb.TaskResult) []pb.TaskResult {
	values := make([]pb.TaskResult, len(taskResults))
	for i, task := range taskResults {
		if task != nil {
			values[i] = *task
		}
	}
	return values
}

func (s *GrpcServer) Execute(ctx context.Context, req *pb.StegAnalysisRequest) (res *pb.StegAnalysisResponse, err error) {
	defer func(start time.Time) {
		duration := time.Since(start).Milliseconds()
		if err != nil {
			StdErr(ctx, "ExecuteWorkflow", err)
		} else {
			log.Printf("executed analysis in %v \n", duration)
		}
		if res == nil {
			res = &pb.StegAnalysisResponse{}
		}
		res.DurationMs = duration
		res.Sha256 = ComputeSHA256(req.File)

		s.rptSvc.ch <- &RptItem{fileName: req.FileName, taskTestults: convertTaskResultPointersToValues(res.TaskResults), durationMs: res.DurationMs, error: res.Error, sha256: res.Sha256}

		filteredRes := make([]*pb.TaskResult, 0)

		for i, v := range res.TaskResults {
			if v == nil {
				StdErr(ctx, "ExecuteWorkflow", fmt.Errorf("error: v is nil!"))
				continue
			}
			task, err := s.stegaSvc.GetTaskByTaskId(v.TaskId)
			if err != nil {
				StdErr(ctx, "ExecuteWorkflow", err)
			} else {
				if !s.stegaSvc.HideOutput() || (task.ShowOutput && !(task.HideResultOnErr && v.Status != string(pb.Status_SUCCESS))) {
					filteredRes = append(filteredRes, res.TaskResults[i])
				}
			}
		}
		res.TaskResults = filteredRes
	}(time.Now())
	s.reqCnt++
	log.Printf("Request: %v", s.reqCnt)
	return s.stegaSvc.Execute(ctx, req)
}

func (s *GrpcServer) GetWorkflowFile(ctx context.Context, req *emptypb.Empty) (*pb.GetWorkflowResponse, error) {
	return s.confSvc.GetWorkflow(ctx, req)
}

func (s *GrpcServer) SetWorkflowFile(ctx context.Context, req *pb.SetWorkflowRequest) (res *pb.SetWorkflowResponse, err error) {
	defer func() {
		if err == nil {
			err = s.stegaSvc.LoadWorkflowConfig()
		}
	}()
	return s.confSvc.SetWorkflow(ctx, req)
}

func (s *GrpcServer) GetServices(ctx context.Context, req *emptypb.Empty) (*pb.GetServicesResponse, error) {
	return s.confSvc.GetServices(ctx, req)
}

func (s *GrpcServer) Start() error {
	listener, err := net.Listen("tcp", s.ctx.gwAddr)
	if err != nil {
		return err
	}
	const MAX_MESSAGE_SIZE = 40 * 1024 * 1024
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(MAX_MESSAGE_SIZE),
		grpc.MaxSendMsgSize(MAX_MESSAGE_SIZE),
	)

	pb.RegisterStegAnalysisServiceServer(server, s)
	pb.RegisterConfigServiceServer(server, s)

	log.Printf("server listening at %v", listener.Addr())

	if err = server.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *GrpcServer) Stop() {
	s.rptSvc.Stop()
}
