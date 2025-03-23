package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	pb "restgw/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestServerContext struct {
	GRPCGateway string
	Addr        string
}

type RestServer struct {
	Context        *RestServerContext
	GRPCSaClient   pb.StegAnalysisServiceClient
	GRPCConfClient pb.ConfigServiceClient
	Mux            *http.ServeMux
	SvcsInfos      []pb.StegServiceInfo
}

func NewRestServerContext() *RestServerContext {
	addr := os.Getenv("port")
	if !strings.Contains(addr, ":") {
		addr = ":" + addr
	}
	return &RestServerContext{
		GRPCGateway: os.Getenv("grpcgw"),
		Addr:        addr,
	}
}

func NewRestServer(ctx *RestServerContext) *RestServer {
	conn, err := grpc.NewClient(ctx.GRPCGateway, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to %v: %v", ctx.GRPCGateway, err)
	}

	return &RestServer{
		Context:        ctx,
		GRPCSaClient:   pb.NewStegAnalysisServiceClient(conn),
		GRPCConfClient: pb.NewConfigServiceClient(conn),
		Mux:            http.NewServeMux(),
	}
}

func (server *RestServer) InitializeRoutes() {
	server.Mux.HandleFunc("/execute", handleStegAnalysisRequest(server.GRPCSaClient))
	server.Mux.HandleFunc("/workflow", handleWorkflowConfigRequest(server.GRPCConfClient))
	server.Mux.HandleFunc("/services", handleGetServicesRequest(server.GRPCConfClient))
	fs := http.FileServer(http.Dir("../clients"))
	server.Mux.Handle("/clients/", http.StripPrefix("/clients", fs))
}

func (server *RestServer) Start() {
	log.Printf("starting server on %v", server.Context.Addr)
	handler := loggingMiddleware(server.Mux)
	handler = corsMiddleware(handler)
	log.Fatal(http.ListenAndServe(server.Context.Addr, handler))
}

func (server *RestServer) GetSvcsInfos() {
	// if server.SvcsInfos == nil {
	// 	server.GRPCConfClient.GetServices()
	// }
}

func getClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}
	ip := r.RemoteAddr
	ip = strings.Split(ip, ":")[0]
	return ip
}
