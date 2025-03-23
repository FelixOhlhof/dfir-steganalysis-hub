package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	pb "grpcgw/pb"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	"google.golang.org/protobuf/types/known/structpb"
)

func copyFileToFolder(sourceFilePath, targetFolderPath string) error {
	targetFilePath := filepath.Join(targetFolderPath, filepath.Base(sourceFilePath))

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		return fmt.Errorf("error creating target file: %w", err)
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return nil
}

func ComputeSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func containsIgnoreCase(slice []string, target string) bool {
	target = strings.ToLower(target)
	for _, str := range slice {
		if strings.ToLower(str) == target {
			return true
		}
	}
	return false
}

func checkFileTypeAllowed(data []byte, allwdFlTypes []string) (bool, string) {
	// enough space
	if len(data) < 4 {
		return true, "unknown"
	}

	switch {
	// PNG: 89 50 4E 47
	case data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47:
		return containsIgnoreCase(allwdFlTypes, "PNG"), "PNG"

	// JPG: FF D8 FF
	case len(data) > 2 && data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF:
		return containsIgnoreCase(allwdFlTypes, "JPG") || containsIgnoreCase(allwdFlTypes, "JPEG"), "JPG"

	// BMP: 42 4D
	case data[0] == 0x42 && data[1] == 0x4D:
		return containsIgnoreCase(allwdFlTypes, "BMP"), "BMP"

	// GIF: 47 49 46 38 (GIF87a, GIF89a)
	case data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x38:
		return containsIgnoreCase(allwdFlTypes, "GIF"), "GIF"

	// TIF: 49 49 2A 00 oder 4D 4D 00 2A
	case (data[0] == 0x49 && data[1] == 0x49 && data[2] == 0x2A && data[3] == 0x00) ||
		(data[0] == 0x4D && data[1] == 0x4D && data[2] == 0x00 && data[3] == 0x2A):
		return containsIgnoreCase(allwdFlTypes, "TIF") || containsIgnoreCase(allwdFlTypes, "TIFF"), "TIF"

	// SVG: XML-Header
	case len(data) > 5 && data[0] == 0x3C && data[1] == 0x3F && data[2] == 0x78 && data[3] == 0x6D && data[4] == 0x6C:
		return containsIgnoreCase(allwdFlTypes, "SVG"), "SVG"

	// ICO: 00 00 01 00
	case data[0] == 0x00 && data[1] == 0x00 && data[2] == 0x01 && data[3] == 0x00:
		return containsIgnoreCase(allwdFlTypes, "ICO"), "ICO"

	// WEBP: 52 49 46 46
	case len(data) > 11 && data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46 &&
		data[8] == 0x57 && data[9] == 0x45 && data[10] == 0x42 && data[11] == 0x50:
		return containsIgnoreCase(allwdFlTypes, "WEBP"), "WEBP"

	default:
		return false, "unknown"
	}
}

func getGoValuateFuncs() map[string]govaluate.ExpressionFunction {
	functions := make(map[string]govaluate.ExpressionFunction)
	functions["safe"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("safe expects exactly one argument")
		}

		if args[0] == nil {
			return args[1], nil
		}

		return args[0], nil
	}
	functions["strlen"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("strlen expects exactly one argument")
		}

		str, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		length := len(str)
		return length, nil
	}
	functions["strcontains"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("strcontains expects exactly one argument")
		}

		str, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		substr, err := parseToString(args[1])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return strings.Contains(str, substr), nil
	}
	functions["strstartswith"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("strstartswith expects exactly one argument")
		}

		str, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		prefix, err := parseToString(args[1])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return strings.HasPrefix(str, prefix), nil
	}
	functions["strendswith"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("strendswith expects exactly one argument")
		}

		str, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		suffix, err := parseToString(args[1])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return strings.HasSuffix(str, suffix), nil
	}
	functions["strtolower"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("strtolower expects exactly one argument")
		}

		str, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return strings.ToLower(str), nil
	}
	functions["strreplace"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("strreplace expects exactly one argument")
		}

		str, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		old, err := parseToString(args[1])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		new, err := parseToString(args[2])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return strings.ReplaceAll(str, old, new), nil
	}
	functions["strsplit"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("strsplit expects exactly one argument")
		}

		cleanStr, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		delimiter, err := parseToString(args[1])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return strings.Split(cleanStr, delimiter), nil
	}
	functions["strtrim"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("strtrim expects exactly one argument")
		}

		cleanStr, err := parseToString(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return strings.TrimSpace(cleanStr), nil
	}
	functions["toString"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("toString expects exactly one argument")
		}

		cleanArg, err := convertToSaveValue(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return fmt.Sprintf("%v", cleanArg), nil
	}
	functions["toNumber"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("toNumber expects exactly one argument")
		}

		cleanArg, err := convertToSaveValue(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		switch v := cleanArg.(type) {
		case string:
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("cannot convert %s to a number", v)
			}
			return f, nil
		case int:
			return float64(v), nil
		case float64:
			return v, nil
		default:
			return nil, fmt.Errorf("unsupported type for toNumber: %T", v)
		}
	}
	functions["toBool"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("toBool expects exactly one argument")
		}

		cleanArg, err := convertToSaveValue(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		switch v := cleanArg.(type) {
		case string:
			b, err := strconv.ParseBool(v)
			if err != nil {
				return nil, fmt.Errorf("cannot convert %s to bool", v)
			}
			return b, nil
		case int:
			return v != 0, nil
		case float64:
			return v != 0.0, nil
		default:
			return nil, fmt.Errorf("unsupported type for toBool: %T", v)
		}
	}
	functions["isNull"] = func(args ...interface{}) (interface{}, error) {
		if len(args) == 0 {
			return true, nil
		}

		cleanArg, err := convertToSaveValue(args[0])
		if err != nil {
			return nil, fmt.Errorf("error converting argument: %v", err)
		}
		return cleanArg == nil, nil
	}
	functions["isIterable"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("isIterable expects exactly one argument")
		}

		arg := args[0]
		val := reflect.ValueOf(arg)

		switch val.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map:
			return true, nil
		default:
			return false, nil
		}
	}
	functions["count"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("isIterable expects exactly one argument")
		}

		strctData, ok := args[0].(*pb.ResponseValue_StructuredValue)
		if ok {
			lst := strctData.StructuredValue.GetListValue()
			dict := strctData.StructuredValue.GetStructValue()

			if lst != nil {
				return len(lst.Values), nil
			}
			if dict != nil {
				return len(dict.GetFields()), nil
			}
		}
		return nil, fmt.Errorf("type not supported")
	}
	functions["isNumeric"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("isNull expects exactly one argument")
		}

		switch args[0].(type) {
		case int, float64, float32:
			return true, nil
		default:
			return false, nil
		}
	}
	functions["sizeOf"] = func(args ...interface{}) (interface{}, error) {
		if len(args) == 0 { // passed nil value
			return 0.0, nil
		}

		if len(args) != 1 {
			return nil, fmt.Errorf("sizeOf expects exactly one argument")
		}

		target := args[0]

		switch value := target.(type) {
		case []uint8:
			return float64(len(value)), nil
		case string:
			return float64(len(value)), nil
		default:
			return 0.0, fmt.Errorf("type %v not supported for funtion sizeOf", value)
		}

		// if target == nil {
		// 	return 0.0, nil
		// }

		// var buf bytes.Buffer
		// encoder := gob.NewEncoder(&buf)
		// if err := encoder.Encode(target); err != nil {
		// 	return 0.0, nil
		// }

		// return float64(buf.Len()), nil
	}
	functions["listContains"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("wrong number of arguments provided")
		}

		target := args[1]

		strctData, ok := args[0].([]*structpb.Value)
		if ok {
			for _, value := range strctData {
				if isValueEqualToInterface(value, target) {
					return true, nil
				}
			}
			return false, nil
		}

		lstData, ok := args[0].([]interface{})
		if ok {
			for _, value := range lstData {
				if value == target {
					return true, nil
				}
			}
			return false, nil
		}

		lstStrData, ok := args[0].([]string)
		if ok {
			for _, value := range lstStrData {
				if value == target {
					return true, nil
				}
			}
			return false, nil
		}

		lstIntData, ok := args[0].([]int)
		if ok {
			for _, value := range lstIntData {
				if value == target {
					return true, nil
				}
			}
			return false, nil
		}

		lstFloat32Data, ok := args[0].([]float32)
		if ok {
			for _, value := range lstFloat32Data {
				if value == target {
					return true, nil
				}
			}
			return false, nil
		}

		lstFloat64Data, ok := args[0].([]float64)
		if ok {
			for _, value := range lstFloat64Data {
				if value == target {
					return true, nil
				}
			}
			return false, nil
		}

		lstBoolData, ok := args[0].([]bool)
		if ok {
			for _, value := range lstBoolData {
				if value == target {
					return true, nil
				}
			}
			return false, nil
		}

		return nil, fmt.Errorf("args[0] is not a list")
	}
	functions["containsKey"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("containsKey requires exactly 2 arguments")
		}

		dict, ok := args[0].(map[string]*structpb.Value)
		if !ok {
			return nil, fmt.Errorf("first argument must be a map[string]interface{}")
		}

		key, err := parseToString(args[1])
		if err != nil {
			return nil, fmt.Errorf("second argument must be a string (the key): %v", err)
		}

		_, exists := dict[key]
		return exists, nil
	}
	functions["condReturn"] = func(args ...interface{}) (interface{}, error) {
		if len(args)%2 != 1 {
			return nil, fmt.Errorf("condReturn requires 2x+1 arguments")
		}

		for i := 0; i < len(args)-1; i += 2 {
			result, ok := args[i].(bool)
			if !ok {
				return nil, fmt.Errorf("argument must be a expresion or a boolean")
			}
			if result {
				return args[i+1], nil
			}
		}
		return args[len(args)-1], nil
	}
	return functions
}

func parseGoEvalExpr(expr string, param map[string]interface{}, funcs map[string]govaluate.ExpressionFunction) (interface{}, error) {
	re := regexp.MustCompile(`([\w.]+->[\w.]+(?:\[\d+\])?(?:\.[\w.]+(?:\[\d+\])?)*)`)
	parsed := re.ReplaceAllStringFunc(expr, func(m string) string {
		return applyGoEvalEscCar(m)
	})

	// add nil values
	matches := re.FindAllStringSubmatch(expr, -1)
	for _, m := range matches {
		if len(m) >= 2 {
			_, ok := param[m[1]]
			if !ok {
				param[m[1]] = nil
			}
		}
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(parsed, funcs)
	if err != nil {
		return nil, fmt.Errorf("error initializing govaluate expression: %v", err)
	}

	// evaluate expression
	result, err := expression.Evaluate(param)
	if err != nil {
		return fmt.Errorf("error evaluating expression: %v", err).Error(), nil
	}
	return result, nil
}

func applyGoEvalEscCar(val string) string {
	replacedEscCaracs := strings.ReplaceAll(val, "[", "\\[")
	replacedEscCaracs = strings.ReplaceAll(replacedEscCaracs, "]", "\\]")
	return fmt.Sprintf("[%v]", replacedEscCaracs)
}

func isValueEqualToInterface(value *structpb.Value, target interface{}) bool {
	if value == nil || target == nil {
		return value == nil && target == nil
	}

	switch v := value.Kind.(type) {
	case *structpb.Value_StringValue:
		strTarget, ok := target.(string)
		return ok && v.StringValue == strTarget
	case *structpb.Value_NumberValue:
		floatTarget, ok := target.(float64)
		return ok && v.NumberValue == floatTarget
	case *structpb.Value_BoolValue:
		boolTarget, ok := target.(bool)
		return ok && v.BoolValue == boolTarget
	case *structpb.Value_NullValue:
		return target == nil
	case *structpb.Value_ListValue:
		listTarget, ok := target.([]interface{})
		if !ok || len(v.ListValue.Values) != len(listTarget) {
			return false
		}
		for i, item := range v.ListValue.Values {
			if !isValueEqualToInterface(item, listTarget[i]) {
				return false
			}
		}
		return true
	case *structpb.Value_StructValue:
		mapTarget, ok := target.(map[string]interface{})
		if !ok || len(v.StructValue.Fields) != len(mapTarget) {
			return false
		}
		for key, field := range v.StructValue.Fields {
			if !isValueEqualToInterface(field, mapTarget[key]) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func convertToSaveValue(val interface{}) (interface{}, error) {
	if spb, ok := val.(*structpb.Value); ok {
		switch kind := spb.Kind.(type) {
		case *structpb.Value_StringValue:
			return kind.StringValue, nil
		case *structpb.Value_NumberValue:
			return kind.NumberValue, nil
		case *structpb.Value_BoolValue:
			return kind.BoolValue, nil
		case *structpb.Value_NullValue:
			return nil, nil
		default:
			return nil, fmt.Errorf("unsupported structpb.Value type: %T", kind)
		}
	}
	return val, nil
}

func parseToString(arg interface{}) (string, error) {
	cleanArg, err := convertToSaveValue(arg)
	if err != nil {
		return "", err
	}
	str, ok := cleanArg.(string)
	if !ok {
		return "", fmt.Errorf("argument can not be converted to a string")
	}
	return str, nil
}
