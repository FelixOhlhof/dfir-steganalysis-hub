package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "restgw/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

func handleGetServicesRequest(client pb.ConfigServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		resp, err := client.GetServices(context.Background(), &emptypb.Empty{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Error converting response to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}

func handleWorkflowConfigRequest(client pb.ConfigServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var upload struct {
				File string `json:"file"`
			}

			err := json.NewDecoder(r.Body).Decode(&upload)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			req := &pb.SetWorkflowRequest{
				Content: upload.File,
			}

			resp, err := client.SetWorkflowFile(context.Background(), req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte(resp.Message))
		} else if r.Method == http.MethodGet {
			resp, err := client.GetWorkflowFile(context.Background(), &emptypb.Empty{})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(resp.Content))
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
	}
}
