package lacia

import (
	"fmt"
	"testing"
)

func TestSortArr(t *testing.T) {
	var sortArray = []int{11, 3, 6, 25, 15, 1, 27}
	fmt.Println("before sort:", sortArray)
	SortArrAsc(sortArray)
	fmt.Println("after sort:", sortArray)

	fmt.Println("before sort:", sortArray)
	SortArrDesc(sortArray)
	fmt.Println("after sort:", sortArray)
}

func TestRepeatElementSidesInt64(t *testing.T) {
	var arr = []int64{11, 15, 6, 11, 15, 1, 15, 7, 88}
	fmt.Println("before:", arr)
	arr0, m := RepeatElementSidesInt64(arr)
	fmt.Println("arr0:", arr0)
	fmt.Println("m:", m)
}

func TestRepeatElementSidesString(t *testing.T) {
	var arr = []string{"F", "V", "R", "V", "A", "P", "F", "A", "F"}
	fmt.Println("before:", arr)
	arr0, m := RepeatElementSidesString(arr)
	fmt.Println("arr0:", arr0)
	fmt.Println("m:", m)
}

func TestExistsInListString(t *testing.T) {
	var arr0 = []string{TimeTemplate1, TimeTemplate2, TimeTemplate3, TimeTemplate5, TimeTemplate6, TimeTemplate1}
	indexS := ExistsInListString(arr0, TimeTemplate3, true)
	fmt.Println(indexS) // [2]
	indexS = ExistsInListString(arr0, TimeTemplate1, true)
	fmt.Println(indexS) // [0]
	indexS = ExistsInListString(arr0, TimeTemplate5, false)
	fmt.Println(indexS) // [3]
	indexS = ExistsInListString(arr0, TimeTemplate1, false)
	fmt.Println(indexS) // [0 5]
	indexS = ExistsInListString(arr0, "", false)
	fmt.Println(indexS) // [-1]
	indexS = ExistsInListString(arr0, "2006", false)
	fmt.Println(indexS) // [-1]
}

func TestExistsInListInt(t *testing.T) {
	// 有序
	var arr0 = []int{1, 2, 3, 4, 5}
	fmt.Println(ExistsInSortList(arr0, 1))
	fmt.Println(ExistsInSortList(arr0, 3))
	fmt.Println(ExistsInSortList(arr0, 5))
	// 无序
	var arr = []int{11, 15, 6, 11, 15, 18, 6, 1, 11, 88, 88}
	fmt.Println("11:", ExistsInListInt(arr, 11, true))  // 11: [0]
	fmt.Println("11:", ExistsInListInt(arr, 11, false)) // 11: [0 8 3]
	fmt.Println("15:", ExistsInListInt(arr, 15, false)) // 15: [1 4]
	fmt.Println("6:", ExistsInListInt(arr, 6, false))   // 6: [2 6]
	fmt.Println("18:", ExistsInListInt(arr, 18, false)) // 18: [5]
	fmt.Println("7:", ExistsInListInt(arr, 7, false))   // 7: [-1]
	fmt.Println("88:", ExistsInListInt(arr, 88, false)) // 88: [10 9]
	fmt.Println("88:", ExistsInListInt(arr, 88, true))  // 88: [10]
	fmt.Println("89:", ExistsInListInt(arr, 89, false)) // 89: [-1]
}
