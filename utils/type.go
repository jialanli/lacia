package lacia

import "fmt"

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
