package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"

	pb "vt/pb"

	vt "github.com/VirusTotal/vt-go"
)

type VtServiceServer struct {
	pb.UnimplementedStegServiceServer
	vtClient *vt.Client
}

func NewVtServiceServer(apiKey string) *VtServiceServer {
	client := vt.NewClient(apiKey)
	return &VtServiceServer{vtClient: client}
}

func (s *VtServiceServer) Execute(ctx context.Context, req *pb.StegServiceRequest) (*pb.StegServiceResponse, error) {
	function := req.GetFunction()
	params := req.GetParams()
	fileData := req.GetFile()

	switch function {
	case "scan_file":
		waitForRep := params["delay"].GetIntValue()
		if waitForRep == 0 {
			waitForRep = 5
		}
		return s.ScanFile(ctx, fileData, time.Duration(waitForRep)*time.Second)
	case "get_report":
		return s.GetFileReport(ctx, fileData)
	default:
		response := &pb.StegServiceResponse{
			Values: make(map[string]*pb.ResponseValue),
		}
		response.Error = "Unsupported function: " + function
		return response, nil
	}
}

func (s *VtServiceServer) GetStegServiceInfo(ctx context.Context, _ *emptypb.Empty) (*pb.StegServiceInfo, error) {
	return &pb.StegServiceInfo{
		Name:        "vt",
		Description: "A service for VirusTotal file analysis and other utilities.",
		Functions: []*pb.StegServiceFunction{
			{
				Name:        "scan_file",
				Description: "Scans a file using the VirusTotal API.",
				Parameter: []*pb.StegServiceParameterDefinition{
					{Name: "delay", Type: pb.Type_INT, Default: "5", Description: "Delay for the scan (in sec)."},
				},
				ReturnFields: []*pb.StegServiceReturnFieldDefinition{
					{Name: "report", Type: pb.Type_DICT, Description: "Report from VirusTotal"},
					{Name: "report_url", Type: pb.Type_STRING, Description: "URL to Analysis."},
				},
			},
			{
				Name:        "get_report",
				Description: "Retrieves the analysis report for a file.",
				ReturnFields: []*pb.StegServiceReturnFieldDefinition{
					{Name: "report", Type: pb.Type_DICT, Description: "Report from VirusTotal"},
					{Name: "report_url", Type: pb.Type_STRING, Description: "URL to Analysis."},
				},
			},
		},
	}, nil
}

func (s *VtServiceServer) ScanFile(ctx context.Context, file []byte, delay time.Duration) (*pb.StegServiceResponse, error) {
	rpt, err := s.GetFileReport(ctx, file)
	if err == nil {
		return rpt, nil
	}

	_, err = s.vtClient.NewFileScanner().Scan(bytes.NewReader(file), "", nil)
	if err != nil {
		return nil, err
	}

	time.Sleep(delay)

	return s.GetFileReport(ctx, file)
}

func (s *VtServiceServer) GetFileReport(_ context.Context, file []byte) (*pb.StegServiceResponse, error) {
	fileHash := fmt.Sprintf("%x", sha256.Sum256(file))

	response := &pb.StegServiceResponse{
		Values: make(map[string]*pb.ResponseValue),
	}

	vtObj, err := s.vtClient.GetObject(vt.URL("files/%s", fileHash))
	if err != nil {
		return nil, err
	}

	report := make(map[string]interface{})

	for _, s := range vtObj.Attributes() {
		report[s], _ = vtObj.Get(s)
	}

	structuredReport, err := structpb.NewStruct(report)
	if err != nil {
		return nil, fmt.Errorf("failed to convert report to protobuf struct: %v", err)
	}

	structuredReportValue := structpb.NewStructValue(structuredReport)

	response.Values["report_url"] = &pb.ResponseValue{Value: &pb.ResponseValue_StringValue{StringValue: fmt.Sprintf("https://www.virustotal.com/gui/file/%s", fileHash)}}

	response.Values["report"] = &pb.ResponseValue{
		Value: &pb.ResponseValue_StructuredValue{
			StructuredValue: structuredReportValue,
		},
	}

	return response, nil
}

func main() {
	port := os.Getenv("port")
	if !strings.Contains(port, ":") {
		port = fmt.Sprintf(":%v", port)
	}

	apiKey := os.Getenv("api_key")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	stegServer := NewVtServiceServer(apiKey)
	pb.RegisterStegServiceServer(grpcServer, stegServer)

	log.Printf("Server is listening on port %v", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
