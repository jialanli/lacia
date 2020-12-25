package lacia

import "fmt"

// 获取类型  --sample
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
	case []string:
		return "[]string"
	case []int:
		return "[]int"
	case []int64:
		return "[]int64"
	case []int8:
		return "[]int8"
	case []int32:
		return "[]int32"
	default:
		return "unknown"
	}
}

func GetValTypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
