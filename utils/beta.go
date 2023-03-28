package lacia

import "strconv"

// new many slice
func NewSlice(valT interface{}, n int) {
	t := GetValTypeOf(valT)
	type sType interface{}
	switch t {
	case "func":
		//sType = sType.(func())
	case "bool":
		//sType = bool(sType)
	case "float64":
		//sType = sType.(bool)
		return
	case "float32":
		return
	case "string":
		return
	case "[]string":
		return
	case "[]int":
		return
	case "[]int64":
		return
	case "[]int8":
		return
	case "[]int32":
		return
	default:

	}
	//var slist sType
	for i := 1; i <= n; i++ {

	}
}

// ***************************日期比较start***************************
/*
传入两个字符串类型的日期,格式可以是YYYYMMDD、YYYYMMDD HH:mm:ss、YYYY-MM-DD、YYYY/MM/DD之一
不论是什么格式，t1和t2两者格式必须相同
return:
	t1=t2  		return 0
	t1>t2  		return 1
	t1<t2  		return 2
	not compare return -1
*/
func CompareTwoStrTime(t1, t2 string) int {
	if len(t1) != len(t2) {
		return -1
	}

	t1, t2 = RemoveX(t1, ` `), RemoveX(t2, ` `)
	count := len(t1)
	for i := 0; i < count; i++ {
		if t1[i] == t2[i] {
			continue
		}
		if !IsNum(string(t1[i])) || !IsNum(string(t2[i])) {
			continue
		}
		n1, err := strconv.Atoi(string(t1[i]))
		if err != nil {
			return -1
		}
		n2, err := strconv.Atoi(string(t2[i]))
		if err != nil {
			return -1
		}

		if n1 > n2 {
			return 1
		}
		if n2 > n1 {
			return 2
		}
	}

	return 0
}
