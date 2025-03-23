package main

import (
	"context"
	pb "grpcgw/pb"
	"io"
	"os"
	"path/filepath"

	"log"
	"time"

	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

var LogFile *os.File

type WfLoggingService struct {
	next IStegAnalysisService
}

type ConfLoggingService struct {
	next IConfigService
}

func StdOutput(ctx context.Context, fn string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		clientIP := p.Addr.String()
		log.Printf("[%v] Request: %s", clientIP, fn)
	} else {
		log.Printf("Request: %s", fn)
	}
}

func StdErr(ctx context.Context, fn string, err error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		clientIP := p.Addr.String()
		log.Printf("[%v] Error executing %s: %s", clientIP, fn, err.Error())
	} else {
		log.Printf("Error executing %s: %s", fn, err.Error())
	}
}

func NewStegLoggingService(next IStegAnalysisService) IStegAnalysisService {
	InitializeLogger()
	return &WfLoggingService{next: next}
}

func NewConfLoggingService(next IConfigService) IConfigService {
	return &ConfLoggingService{next: next}
}

func InitializeLogger() {
	if LogFile != nil {
		LogFile.Close()
	}

	LogFile, _ = os.OpenFile(filepath.Join(OutputDir, "results.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	multiWriter := io.MultiWriter(os.Stdout, LogFile)
	log.SetOutput(multiWriter)
}

func (s *WfLoggingService) Execute(ctx context.Context, req *pb.StegAnalysisRequest) (res *pb.StegAnalysisResponse, err error) {
	defer func(start time.Time) {
		duration := time.Since(start).Milliseconds()
		if err != nil {
			StdErr(ctx, "ExecuteWorkflow", err)
		} else {
			log.Printf("executed analysis in %v \n", duration)
		}
		res.DurationMs = duration
	}(time.Now())

	StdOutput(ctx, "ExecuteWorkflow")

	return s.next.Execute(ctx, req)
}

func (svc *WfLoggingService) LoadWorkflowConfig() error {
	defer log.Print("loaded workflow")
	return svc.next.LoadWorkflowConfig()
}

func (s *ConfLoggingService) GetWorkflow(ctx context.Context, a *emptypb.Empty) (res *pb.GetWorkflowResponse, err error) {
	defer func() {
		if err != nil {
			StdErr(ctx, "GetWorkflow", err)
		}
	}()

	StdOutput(ctx, "GetWorkflow")

	return s.next.GetWorkflow(ctx, a)
}

func (s *ConfLoggingService) SetWorkflow(ctx context.Context, req *pb.SetWorkflowRequest) (res *pb.SetWorkflowResponse, err error) {
	defer func() {
		if err != nil {
			StdErr(ctx, "SetWorkflow", err)
		} else {
			InitializeLogger()
		}
	}()

	StdOutput(ctx, "SetWorkflow")

	return s.next.SetWorkflow(ctx, req)
}

func (s *ConfLoggingService) GetServices(ctx context.Context, a *emptypb.Empty) (res *pb.GetServicesResponse, err error) {
	defer func() {
		if err != nil {
			StdErr(ctx, "GetServices", err)
		}
	}()

	StdOutput(ctx, "GetServices")

	return s.next.GetServices(ctx, a)
}

func (s *WfLoggingService) GetTaskByTaskId(taskId string) (Task, error) {
	return s.next.GetTaskByTaskId(taskId)
}

func (s *WfLoggingService) HideOutput() bool {
	return s.next.HideOutput()
}
