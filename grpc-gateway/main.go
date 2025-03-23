package main

import (
	"log"
)

const ReportCsvName = "results.csv"
const WfFileName = "workflow.yaml"
const OutputDirName = "results"

var WfFileHash uint32

func main() {
	// create grpc api server
	grpcSvr := NewGrpcServer()

	// start grpc api server
	if err := grpcSvr.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

	grpcSvr.Stop()
	log.Printf("shutting down...")
}
