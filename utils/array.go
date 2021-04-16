package lacia

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

// 排序:sortBy传空时默认按降序,传asc时按升序
func SortArrAsc(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	base := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[i] > base {
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

	SortArrAsc(arr[:left])
	SortArrAsc(arr[left+1:])
	return arr
}

func SortArrDesc(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	base := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
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

	SortArrDesc(arr[:left])
	SortArrDesc(arr[left+1:])
	return arr
}

func SortFloat64ArrAsc(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}
	base := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[i] > base {
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

	SortFloat64ArrAsc(arr[:left])
	SortFloat64ArrAsc(arr[left+1:])
	return arr
}

func SortFloat64ArrDesc(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}
	base := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
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

	SortFloat64ArrDesc(arr[:left])
	SortFloat64ArrDesc(arr[left+1:])
	return arr
}

func SearchBy2Sides(arr []int, num int) int {
	var low, high int
	for low <= high {
		midIndex := (low + high) / 2
		if arr[midIndex] == num {
			return midIndex
		} else if num < arr[midIndex] {
			high = midIndex - 1
		} else if num > arr[midIndex] {
			low = midIndex + 1
		}
	}

	return -1
}

// 判断有序[]中是否存在某元素,存在则返回索引,不存在则返回-1
// num为要判定是否存在的元素
func ExistsInSortList(list []int, num int) (index int) {
	func(arr []int, low, high, num int) {
		for low <= high {
			midIndex := (low + high) / 2
			if arr[midIndex] == num {
				index = midIndex
				return
			} else if num < arr[midIndex] {
				high = midIndex - 1
			} else if num > arr[midIndex] {
				low = midIndex + 1
			}
		}

		index = -1
		return
	}(list, 0, len(list)-1, num)
	return
}

// 判断[]中是否存在某元素,存在则返回出现该元素的所有索引列表,不存在则返回indexList[0]=-1
// str为要判定是否存在的元素,onlyExists为只需判断存在标志位，若onlyExists=true,则判断到存在num时即返回(indexList[0]不为-1)，否则持续判断收集出现该元素的所有索引列表
// 传false时也相当于统计str在list中出现的次数
func ExistsInListString(list []string, str string, onlyExists bool) (indexList []int) {
	if len(list) <= 0 {
		return nil
	}
	x, y := 0, len(list)-1
	func(arr []string, element string) {
		for y-x >= 0 {
			if element == list[x] {
				indexList = append(indexList, x)
				if onlyExists || x == y {
					return
				}
			}
			x++
			if element == list[y] {
				indexList = append(indexList, y)
				if onlyExists {
					return
				}
			}
			y--
		}

		if len(indexList) < 1 {
			indexList = make([]int, 1)
			indexList[0] = -1
		}
		return
	}(list, str)
	return
}

// 判断[]中是否存在某元素,存在则返回出现该元素的所有索引列表,不存在则返回indexList[0]=-1
// num为要判定是否存在的元素,onlyExists为只需判断存在标志位，若onlyExists=true,则判断到存在num时即返回(indexList[0]不为-1)，否则持续判断收集出现该元素的所有索引列表
func ExistsInListInt(list []int, num int, onlyExists bool) (indexList []int) {
	if len(list) <= 0 {
		return nil
	}
	x, y := 0, len(list)-1
	func(arr []int, element int) {
		for y-x >= 0 {
			if element == list[x] {
				indexList = append(indexList, x)
				if onlyExists || x == y {
					return
				}
			}
			x++
			if element == list[y] {
				indexList = append(indexList, y)
				if onlyExists {
					return
				}
			}
			y--
		}

		if len(indexList) < 1 {
			indexList = make([]int, 1)
			indexList[0] = -1
		}
		return
	}(list, num)
	return
}

type ArrList struct {
	Int64Data   int64
	IntData     int
	Num         int // 最后一次出现的下标
	RepeatCount int // 重复次数
	StringData  string
}

/*
数组去重：给定一个int64列表,返回去重后的元素列表、元素的重复次数
返回值说明
noRepeatList:去重后的元素列表
m:key为每个元素,value.RepeatCount为重复的次数
*/
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

func RepeatElementSidesInt(list []int) ([]int, map[int]ArrList) {
	var noRepeatList []int
	m := make(map[int]ArrList)
	for i, ID := range list {
		obj := ArrList{IntData: ID, Num: i, RepeatCount: 1}
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

/*
数组去重：给定一个string列表,返回去重后的元素列表、元素的重复次数
返回值说明
noRepeatList:去重后的元素列表
m:key为重复的元素,value.RepeatCount为重复的次数
*/
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
