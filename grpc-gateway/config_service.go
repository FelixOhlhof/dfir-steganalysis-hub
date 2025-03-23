package main

import (
	"context"
	"encoding/base64"
	"fmt"
	pb "grpcgw/pb"
	"hash/adler32"
	"os"

	"google.golang.org/protobuf/types/known/emptypb"
)

type IConfigService interface {
	GetWorkflow(context.Context, *emptypb.Empty) (*pb.GetWorkflowResponse, error)
	SetWorkflow(context.Context, *pb.SetWorkflowRequest) (*pb.SetWorkflowResponse, error)
	GetServices(context.Context, *emptypb.Empty) (*pb.GetServicesResponse, error)
}

type ConfigService struct {
	ctx *GrpcServerContext
}

func NewConfigService(ctx *GrpcServerContext) IConfigService {
	return &ConfigService{ctx: ctx}
}

func (s *ConfigService) GetWorkflow(ctx context.Context, _ *emptypb.Empty) (*pb.GetWorkflowResponse, error) {
	content, err := os.ReadFile(WfFileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	base64Data := base64.StdEncoding.EncodeToString(content)

	return &pb.GetWorkflowResponse{
		Content: base64Data,
	}, nil
}

func (s *ConfigService) SetWorkflow(ctx context.Context, req *pb.SetWorkflowRequest) (*pb.SetWorkflowResponse, error) {
	decoded, err := base64.StdEncoding.DecodeString(req.Content)
	if err != nil {
		return nil, fmt.Errorf("error decoding file: %v", err)
	}

	// calc checkSum
	hasher := adler32.New()
	hasher.Write(decoded)
	checkSum := hasher.Sum32()

	if WfFileHash == checkSum {
		return &pb.SetWorkflowResponse{
			Message: "File already up to date",
		}, nil
	}

	err = os.WriteFile(WfFileName, decoded, 0644)
	if err != nil {
		return nil, fmt.Errorf("error decoding file: %v", err)
	}

	// update check sum
	WfFileHash = checkSum

	// create new output dir
	OutputDir, err = createOutputFolder()
	if err != nil {
		return nil, fmt.Errorf("error creating output dir: %v", err)
	}

	return &pb.SetWorkflowResponse{
		Message: "File updated successfully",
	}, nil
}

func (s *ConfigService) GetServices(ctx context.Context, req *emptypb.Empty) (*pb.GetServicesResponse, error) {
	svcs := make([]*pb.StegServiceInfo, 0)
	for _, v := range s.ctx.stegSvcs {
		svcs = append(svcs, v.svcInfo)
	}
	res := pb.GetServicesResponse{Services: svcs}
	return &res, nil
}
