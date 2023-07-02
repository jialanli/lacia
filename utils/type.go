package lacia

import (
	"fmt"
	"reflect"
	"strconv"
)

func SwitchFromString(src string, kind reflect.Kind) (dist interface{}, err error) {
	switch kind {
	case reflect.Int:
		dist, err = strconv.Atoi(src)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dist, err = strconv.ParseInt(src, 10, 64)
	case reflect.Float32:
		var dist64 float64
		dist64, err = strconv.ParseFloat(src, 32)
		dist = float32(dist64)
	case reflect.Float64:
		dist, err = strconv.ParseFloat(src, 64)
	}

	return
}

// --sample
func GetValTypeOfSample(t interface{}) string {
	switch t.(type) {
	case func():
		return "func"
	case bool:
		return "bool"
	case float64:
		return "float64"
	case float32:
		return "float32"
	case string:
		return "float32"
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case []string:
		return "[]string"
	case []int:
		return "[]int"
	case []int8:
		return "[]int8"
	case []int16:
		return "[]int16"
	case []int32:
		return "[]int32"
	case []int64:
		return "[]int64"
	case []uint:
		return "[]uint"
	case []uint8:
		return "[]uint8"
	case []uint16:
		return "[]uint16"
	case []uint32:
		return "[]uint32"
	case []uint64:
		return "[]uint64"
	default:
		return "unknown"
	}
}

func GetValTypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
