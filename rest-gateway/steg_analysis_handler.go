package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	pb "restgw/pb"
)

func handleStegAnalysisRequest(client pb.StegAnalysisServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var analyFile []byte
		var fileName string = ""
		var exec string

		err := r.ParseMultipartForm(50 << 20) // max total size 50 MB
		if err != nil {
			errStr := fmt.Sprintf("error parsing multipart form %s\n", err)
			http.Error(w, errStr, http.StatusBadRequest)
			log.Print(errStr)
			return
		}

		params := make(map[string]string)

		// extract params
		for key, values := range r.Form {
			for _, value := range values {
				// check for optional parameter
				if key == "exec" {
					exec = value
				} else {
					params[key] = value
				}
			}
		}

		// extract target file and additional data parameter
		for key := range r.MultipartForm.File {
			file, fileHeader, err := r.FormFile(key)
			if err != nil {
				errStr := fmt.Sprintf("error in reading the file %s\n", err)
				http.Error(w, errStr, http.StatusBadRequest)
				log.Print(errStr)
				return
			}
			defer file.Close()

			fileContent, err := io.ReadAll(file)
			if err != nil {
				errStr := fmt.Sprintf("error in reading the file buffer %s\n", err)
				http.Error(w, errStr, http.StatusBadRequest)
				log.Print(errStr)
				return
			}

			// check for target file
			if key == "file" {
				analyFile = fileContent
				fileName = fileHeader.Filename
			} else {
				params[key] = string(fileContent[:])
			}
		}

		if analyFile == nil && exec != "" { // allow empty file (some services may have functions that do not require a file)
			errStr := "no file provided\n"
			http.Error(w, errStr, http.StatusBadRequest)
			log.Print(errStr)
			return
		}

		gRPCReq := &pb.StegAnalysisRequest{
			File:     analyFile,
			FileName: fileName,
			Params:   params,
			Exec:     exec,
		}

		gRPCResp, err := client.Execute(context.Background(), gRPCReq)
		if err != nil {
			http.Error(w, "failed to call gRPC service: "+err.Error(), http.StatusInternalServerError)
			log.Print("failed to call gRPC service: " + err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gRPCResp)
	}
}
