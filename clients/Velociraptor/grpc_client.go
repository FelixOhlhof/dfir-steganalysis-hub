package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "grpc-client/pb"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const MAX_FILE_SIZE = 20 * 1024 * 1024 // 20MB
const MAX_PARALLEL_REQUESTS = 10

func handleErr(err interface{}) {
	var errors []string

	switch v := err.(type) {
	case string:
		errors = []string{v}
	case []string:
		errors = v
	default:
		fmt.Println("invalid param: expected []string")
		os.Exit(1)
	}

	errorMap := make([]map[string]string, len(errors))
	for i, e := range errors {
		errorMap[i] = map[string]string{"Client Error": e}
	}

	jsonOutput, _ := json.Marshal(errorMap)
	fmt.Println(string(jsonOutput))
	os.Exit(1)
}

func printHelp() {
	handleErr(`Usage: <url_to_gateway> <path_or_file> [-x file extensions ...] [key value parameter]

Arguments:
  url_to_gateway     URL to the gateway
  path_or_file       Path or filename to process
Optional Flags:
  -x                List of file extensions to include, e.g., -x png jpg txt
  -r 				Search files recursively
Additional Parameters:
  Arbitrary key-value pairs like -key value, e.g., -arbitraryParam1 paramValue1`)
}

func getFilesWithExtensions(root string, extensions []string, recursive bool) ([]string, error) {
	var files []string

	// check file extensions
	hasMatchingExtension := func(filename string, extensions []string) bool {
		if len(extensions) == 0 {
			return true
		}
		for _, ext := range extensions {
			if strings.HasSuffix(filename, ext) {
				return true
			}
		}
		return false
	}

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !recursive && d.IsDir() && path != root {
			return filepath.SkipDir
		}

		if !d.IsDir() && hasMatchingExtension(d.Name(), extensions) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func parse_args() (gwAdr string, files []string, params map[string]string, deb bool) {
	if len(os.Args) < 3 {
		printHelp()
		return
	}

	gwa := os.Args[1]
	path := os.Args[2]
	recursive := false
	debug := false

	actFlag := ""
	p := make(map[string]string, 0)
	fileExt := make([]string, 0)

	for _, v := range os.Args[3:] {
		if strings.HasPrefix(v, "-x ") {
			actFlag = "-x"
			tmp := strings.Split(v, " ")
			if len(tmp) > 1 {
				fileExt = tmp[1:]
			}
		} else if v == "-r" {
			recursive = true
		} else if v == "-d" {
			debug = true
		} else if strings.HasPrefix(v, "-") {
			actFlag = v
			tmp := strings.Split(v, " ")
			if len(tmp) > 1 {
				p[strings.TrimLeft(tmp[0], "-")] = tmp[1]
			} else {
				p[strings.TrimLeft(v, "-")] = ""
			}
		} else if actFlag == "-x" {
			fileExt = append(fileExt, v)
		} else if actFlag != "" {
			p[strings.TrimLeft(actFlag, "-")] = v
		}
	}

	fl := make([]string, 0)

	info, err := os.Stat(path)
	if err != nil {
		handleErr(err.Error())
	}

	if info.IsDir() {
		files, err := getFilesWithExtensions(path, fileExt, recursive)
		if err != nil {
			handleErr(err.Error())
		}
		fl = files
	} else {
		fl = append(fl, path)
	}

	return gwa, fl, p, debug
}

func execute(file string, params map[string]string, client pb.StegAnalysisServiceClient) *pb.StegAnalysisResponse {
	f, err := os.Open(file)
	if err != nil {
		return &pb.StegAnalysisResponse{Error: fmt.Sprintf("unable to open file: %v", err)}
	}
	defer f.Close()
	fStat, err := f.Stat()
	if err != nil {
		return &pb.StegAnalysisResponse{Error: fmt.Sprintf("unable to retrieve size: %v", err)}
	}
	fSz := fStat.Size()
	if fSz > MAX_FILE_SIZE {
		return &pb.StegAnalysisResponse{Error: fmt.Sprintf("file to big: %v", fSz)}
	}
	fData, err := io.ReadAll(f)
	if err != nil {
		return &pb.StegAnalysisResponse{Error: fmt.Sprintf("unable to read file: %v", err)}
	}

	ctx := context.Background()

	res, err := client.Execute(ctx, &pb.StegAnalysisRequest{File: fData, Params: params, FileName: file})
	if err != nil {
		return &pb.StegAnalysisResponse{Error: fmt.Sprintf("failed to send request: %v", err.Error())}
	}

	return res
}

func flattenResponse(fileName string, res *pb.StegAnalysisResponse) map[string]interface{} {
	result := make(map[string]interface{})

	result["File"] = fileName
	result["Sha256"] = res.Sha256
	if res.Error != "" {
		result["Error"] = res.Error
	}

	for _, r := range res.TaskResults {
		if r.ServiceResponse != nil {
			for key, value := range r.ServiceResponse.Values {
				switch vOut := value.Value.(type) {
				case *pb.ResponseValue_StringValue:
					result[fmt.Sprintf("%v %v", r.TaskId, key)] = vOut.StringValue
				case *pb.ResponseValue_IntValue:
					result[fmt.Sprintf("%v %v", r.TaskId, key)] = fmt.Sprintf("%d", vOut.IntValue)
				case *pb.ResponseValue_FloatValue:
					result[fmt.Sprintf("%v %v", r.TaskId, key)] = fmt.Sprintf("%f", vOut.FloatValue)
				case *pb.ResponseValue_BoolValue:
					result[fmt.Sprintf("%v %v", r.TaskId, key)] = fmt.Sprintf("%t", vOut.BoolValue)
				case *pb.ResponseValue_BinaryValue:
					result[fmt.Sprintf("%v %v", r.TaskId, key)] = fmt.Sprintf("[bytes:%d]", len(vOut.BinaryValue))
				}
			}
		}

		if r.Error != "" {
			result[fmt.Sprintf("%v Error", r.TaskId)] = r.Error
		}
	}

	return result
}

func main() {
	gwAdr, files, params, debug := parse_args()

	if debug {
		var result string
		for i, arg := range os.Args[1:] {
			result += fmt.Sprintf("Arg%d=%s ", i+1, arg)
		}
		handleErr(result)
	}

	conn, err := grpc.NewClient(gwAdr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handleErr(fmt.Sprintf("failed to connect to %v: %v", gwAdr, err))
	}
	defer conn.Close()

	client := pb.NewStegAnalysisServiceClient(conn)

	var wg sync.WaitGroup
	sem := make(chan struct{}, MAX_PARALLEL_REQUESTS)
	results := make([]map[string]interface{}, 0)
	mu := sync.Mutex{}

	for _, f := range files {
		wg.Add(1)
		sem <- struct{}{}

		go func(file string) {
			defer wg.Done()
			defer func() { <-sem }()

			res := execute(file, params, client)
			fRes := flattenResponse(filepath.Base(file), res)

			mu.Lock()
			results = append(results, fRes)
			mu.Unlock()
		}(f)
	}

	wg.Wait()

	// for _, f := range files {
	// 	res := execute(f, params, client)
	// 	fRes := flattenResponse(filepath.Base(f), res)
	// 	results = append(results, fRes)
	// }

	js, err := json.Marshal(results)
	if err != nil {
		handleErr(fmt.Sprintf("failed to marshal response: %v", err.Error()))
	}

	fmt.Println(string(js))
}
