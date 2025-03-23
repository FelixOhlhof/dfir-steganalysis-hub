package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	pb "grpcgw/pb"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
)

type RptSvc struct {
	ch   chan *RptItem
	done chan struct{}
	wg   sync.WaitGroup
}

type RptItem struct {
	fileName     string
	sha256       string
	taskTestults []pb.TaskResult
	durationMs   int64
	error        string
}

func NewRptSvc() RptSvc {
	ch := make(chan *RptItem, 100)

	return RptSvc{ch: ch}
}

func (s *RptSvc) Start() {
	s.wg.Add(1)
	go s.consume()
}

func (s *RptSvc) Stop() {
	close(s.done)
	s.wg.Wait()
	close(s.ch)
}

func (s *RptSvc) consume() {
	defer s.wg.Done()
	for {
		select {
		case res := <-s.ch:
			if err := writeCsvReport(OutputDir, res); err != nil {
				fmt.Printf("error writing csv report: %v", err)
				return
			}
		case <-s.done:
			return
		}
	}
}

// Function to update or rewrite a CSV file with a single new row
func updateOrRewriteCSV(filename string, newRow map[string]string, priorityColumns []string) error {
	// Open the file or create a new one if it does not exist
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, create a new one
			return createNewCSV(filename, []map[string]string{newRow}, priorityColumns)
		}
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)
	reader.Comma = ';'
	existingHeader, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	// Generate the updated header order based on priorities and new row
	columnOrder := generateHeaderOrder(existingHeader, []map[string]string{newRow}, priorityColumns)

	// Check if the existing header matches the new header order
	if !headerMatches(existingHeader, columnOrder) {
		fmt.Println("Header mismatch. Rewriting the entire CSV.")
		// Read all existing rows
		existingRows, err := reader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to read existing rows: %w", err)
		}

		// Convert existing rows into map format
		existingRowMaps := rowsToMap(existingRows, existingHeader)

		// Append the new row to existing rows
		allRows := append(existingRowMaps, newRow)

		// Rewrite the entire CSV
		return rewriteCSV(filename, allRows, columnOrder)
	}

	// Append the new row to the existing file
	return appendRowToCSV(filename, newRow, columnOrder)
}

// Function to create a new CSV file with a given set of rows
func createNewCSV(filename string, rows []map[string]string, priorityColumns []string) error {
	// Generate the header order based on priorities and rows
	header := generateHeaderOrder(nil, rows, priorityColumns)

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create new CSV: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	// Write the header
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write the rows
	for _, row := range rows {
		line := make([]string, len(header))
		for i, col := range header {
			line[i] = row[col]
		}
		if err := writer.Write(line); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}
	return nil
}

// Function to rewrite the entire CSV file with new rows and column order
func rewriteCSV(filename string, rows []map[string]string, columnOrder []string) error {
	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to recreate CSV: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	// Write the header
	if err := writer.Write(columnOrder); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write all rows
	for _, row := range rows {
		line := make([]string, len(columnOrder))
		for i, col := range columnOrder {
			line[i] = row[col]
		}
		if err := writer.Write(line); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}

	return nil
}

// Function to append a single new row to an existing CSV file
func appendRowToCSV(filename string, newRow map[string]string, columnOrder []string) error {
	// Open the file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for appending: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	// Write the new row
	line := make([]string, len(columnOrder))
	for i, col := range columnOrder {
		line[i] = newRow[col]
	}
	if err := writer.Write(line); err != nil {
		return fmt.Errorf("failed to write new row: %w", err)
	}

	return nil
}

// Function to generate the header order based on existing headers, new rows, and priority columns
func generateHeaderOrder(existingHeader []string, newRows []map[string]string, priorityColumns []string) []string {
	headerMap := make(map[string]struct{})
	for _, col := range existingHeader {
		headerMap[col] = struct{}{}
	}

	for _, row := range newRows {
		for col := range row {
			headerMap[col] = struct{}{}
		}
	}

	columnOrder := []string{}
	seen := make(map[string]bool)

	for _, col := range priorityColumns {
		if _, exists := headerMap[col]; exists {
			columnOrder = append(columnOrder, col)
			seen[col] = true
		}
	}

	remainingColumns := []string{}
	for col := range headerMap {
		if !seen[col] {
			remainingColumns = append(remainingColumns, col)
		}
	}

	sort.Strings(remainingColumns)

	columnOrder = append(columnOrder, remainingColumns...)

	return columnOrder
}

// Function to convert existing CSV rows into a map
func rowsToMap(rows [][]string, header []string) []map[string]string {
	rowMaps := []map[string]string{}
	for _, row := range rows {
		rowMap := make(map[string]string)
		for i, col := range header {
			if i < len(row) {
				rowMap[col] = row[i]
			}
		}
		rowMaps = append(rowMaps, rowMap)
	}
	return rowMaps
}

// Function to check if two headers match
func headerMatches(existingHeader, newHeader []string) bool {
	if len(existingHeader) != len(newHeader) {
		return false
	}
	for i, col := range existingHeader {
		if col != newHeader[i] {
			return false
		}
	}
	return true
}

func writeCsvReport(rptFolder string, rptItem *RptItem) error {
	newRow := make(map[string]string)
	newRow["sha256"] = rptItem.sha256
	newRow["file_name"] = rptItem.fileName
	newRow["duration"] = strconv.FormatInt(rptItem.durationMs, 10)
	newRow["error"] = rptItem.error

	for _, tskRes := range rptItem.taskTestults {
		newRow[fmt.Sprintf("%v.status", tskRes.TaskId)] = tskRes.Status
		newRow[fmt.Sprintf("%v.duration", tskRes.TaskId)] = strconv.FormatInt(tskRes.DurationMs, 10)

		if tskRes.ServiceResponse != nil {
			for valName, val := range tskRes.ServiceResponse.Values {
				parsedVal, err := getCsvValue(val)

				if err != nil {
					return fmt.Errorf("error parsing response value: %v", err)
				}

				newRow[fmt.Sprintf("%v.%v", tskRes.TaskId, valName)] = parsedVal
			}
		}
		if tskRes.Error != "" {
			newRow[fmt.Sprintf("%v.Error", tskRes.TaskId)] = tskRes.Error
		}
	}

	csvFile := filepath.Join(rptFolder, ReportCsvName)
	prioCols := []string{"sha256", "file_name", "duration", "error"}
	updateOrRewriteCSV(csvFile, newRow, prioCols)

	return nil
}

func getCsvValue(value *pb.ResponseValue) (string, error) {
	if value == nil {
		return "", nil
	}
	switch v := value.Value.(type) {
	case *pb.ResponseValue_StringValue:
		return v.StringValue, nil
	case *pb.ResponseValue_IntValue:
		return strconv.FormatInt(v.IntValue, 10), nil
	case *pb.ResponseValue_FloatValue:
		return strconv.FormatFloat(float64(v.FloatValue), 'f', -1, 32), nil
	case *pb.ResponseValue_BoolValue:
		return strconv.FormatBool(v.BoolValue), nil
	case *pb.ResponseValue_StructuredValue:
		jsonData, err := json.Marshal(v.StructuredValue)
		if err != nil {
			return "", fmt.Errorf("failed to marshal structured value to JSON: %w", err)
		}
		return string(jsonData), nil
	case *pb.ResponseValue_BinaryValue:
		return string(v.BinaryValue), nil
	default:
		return "type not supported", nil
	}
}
