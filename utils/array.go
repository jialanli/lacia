package lacia

import (
	"strings"
)

type ArrList struct {
	Int64Data   int64
	IntData     int
	Num         int // 最后一次出现的下标
	RepeatCount int // 重复次数
	StringData  string
}

// 数组去重算法：给定一个int64列表,得到去重元素列表/找出重复元素
// 一次遍历拿到不重复元素组成的列表(param1)、重复元素(param2)
func RepeatElementSidesInt64(list []int64) ([]int64, map[int64]ArrList) {
	var noRepeatList []int64
	m := make(map[int64]ArrList)
	for i, ID := range list {
		obj := ArrList{Int64Data: ID, Num: i, RepeatCount: 1}
		if data, ok := m[ID]; ok {
			obj.RepeatCount = data.RepeatCount + 1
			m[ID] = obj
			continue
		}
		m[ID] = obj
		noRepeatList = append(noRepeatList, ID)
	}

	return noRepeatList, m
}

// 数组去重算法：给定一个string列表,得到去重元素列表/找出重复元素
func RepeatElementSidesString(list []string) ([]string, map[string]ArrList) {
	var noRepeatList []string
	m := make(map[string]ArrList)
	for i, ID := range list {
		obj := ArrList{StringData: ID, Num: i, RepeatCount: 1}
		if data, ok := m[ID]; ok {
			obj.RepeatCount = data.RepeatCount + 1
			m[ID] = obj
			continue
		}
		m[ID] = obj
		noRepeatList = append(noRepeatList, ID)
	}

	return noRepeatList, m
}

func SortArr(arr []int, sortBy string) []int {
	if len(arr) <= 1 {
		return arr
	}
	base := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if strings.ToUpper(sortBy) == "ASC" {
			if arr[i] > base { // 从小到大
				if i != right {
					arr[i], arr[right] = arr[right], arr[i]
				}
				right--
			} else {
				if i != left {
					arr[i], arr[left] = arr[left], arr[i]
				}
				left++
				i++
			}
		}
		if strings.ToUpper(sortBy) == "DESC" {
			if arr[i] < base {
				if i != right {
					arr[i], arr[right] = arr[right], arr[i]
				}
				right--
			} else {
				if i != left {
					arr[i], arr[left] = arr[left], arr[i]
				}
				left++
				i++
			}
		}
	}

	SortArr(arr[:left], sortBy)
	SortArr(arr[left+1:], sortBy)
	return arr
}

// 判断[]中是否存在某元素
