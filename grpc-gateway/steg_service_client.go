package main

import (
	"context"
	"fmt"
	pb "grpcgw/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// this type contains state of the server
type StegServiceContext struct {
	// client to GRPC service
	client pb.StegServiceClient

	// some other useful objects, like config
	// or logger (to replace global logging)
	// (...)
}

type StegService struct {
	ctx     StegServiceContext
	addr    string
	svcInfo *pb.StegServiceInfo
}

func NewStegServiceClient(addr string) StegService {
	return StegService{addr: addr}
}

func GetServiceByName(services []StegService, name string) (*StegService, error) {
	for i := range services {
		if services[i].svcInfo.Name == name {
			return &services[i], nil
		}
	}
	return nil, fmt.Errorf("no service %v existing", name)
}

func GetFunctionByName(s *StegService, name string) (*pb.StegServiceFunction, error) {
	for i := range s.svcInfo.Functions {
		if s.svcInfo.Functions[i].Name == name {
			return s.svcInfo.Functions[i], nil
		}
	}
	return nil, fmt.Errorf("no function %v existing", name)
}

func (svc *StegService) Connect() error {
	conn, err := grpc.NewClient(svc.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to %v: %v", svc.addr, err)
	}

	svc.ctx = StegServiceContext{
		client: pb.NewStegServiceClient(conn),
	}

	res, err := svc.ctx.client.GetStegServiceInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("failed to retrieve service info from %v: %v", svc.addr, err)
	}
	svc.svcInfo = res

	return nil
}

func (s *StegService) Execute(ctx context.Context, request *pb.StegServiceRequest) (*pb.StegServiceResponse, error) {
	deadline, _ := ctx.Deadline()
	timeout := int64(time.Until(deadline).Seconds())
	if timeout > 0 {
		request.RequestTimeoutSec = timeout
	}
	return s.ctx.client.Execute(ctx, request)
}
