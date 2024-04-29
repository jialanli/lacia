package lacia

import (
	"log"
	"testing"
)

func TestSortArr(t *testing.T) {
	var sortArray = []int{11, 3, 6, 25, 15, 1, 27}
	log.Println("before sort:", sortArray)
	SortArrAsc(sortArray)
	log.Println("after sort:", sortArray)

	log.Println("before sort:", sortArray)
	SortArrDesc(sortArray)
	log.Println("after sort:", sortArray)
}

func TestRepeatElementSidesInt64(t *testing.T) {
	var arr = []int64{11, 15, 6, 11, 15, 1, 15, 7, 88}
	log.Println("before:", arr)
	arr0, m := RepeatElementSidesInt64(arr)
	log.Println("arr0:", arr0)
	log.Println("m:", m)
}

func TestRepeatElementSidesString(t *testing.T) {
	var arr = []string{"F", "V", "R", "V", "A", "P", "F", "A", "F"}
	log.Println("before:", arr)
	arr0, m := RepeatElementSidesString(arr)
	log.Println("arr0:", arr0)
	log.Println("m:", m)
}

func TestExistsInListString(t *testing.T) {
	var arr0 = []string{TimeTemplate1, TimeTemplate2, TimeTemplate3, TimeTemplate5, TimeTemplate6, TimeTemplate1}
	indexS := ExistsInListString(arr0, TimeTemplate3, true)
	log.Println(indexS) // [2]
	indexS = ExistsInListString(arr0, TimeTemplate1, true)
	log.Println(indexS) // [0]
	indexS = ExistsInListString(arr0, TimeTemplate5, false)
	log.Println(indexS) // [3]
	indexS = ExistsInListString(arr0, TimeTemplate1, false)
	log.Println(indexS) // [0 5]
	indexS = ExistsInListString(arr0, "", false)
	log.Println(indexS) // [-1]
	indexS = ExistsInListString(arr0, "2006", false)
	log.Println(indexS) // [-1]
}

func TestExistsInListInt(t *testing.T) {
	// 有序
	var arr0 = []int{1, 2, 3, 4, 5}
	log.Println(ExistsInSortList(arr0, 1))
	log.Println(ExistsInSortList(arr0, 3))
	log.Println(ExistsInSortList(arr0, 5))
	// 无序
	var arr = []int{11, 15, 6, 11, 15, 18, 6, 1, 11, 88, 88}
	log.Println("11:", ExistsInListInt(arr, 11, true))  // 11: [0]
	log.Println("11:", ExistsInListInt(arr, 11, false)) // 11: [0 8 3]
	log.Println("15:", ExistsInListInt(arr, 15, false)) // 15: [1 4]
	log.Println("6:", ExistsInListInt(arr, 6, false))   // 6: [2 6]
	log.Println("18:", ExistsInListInt(arr, 18, false)) // 18: [5]
	log.Println("7:", ExistsInListInt(arr, 7, false))   // 7: [-1]
	log.Println("88:", ExistsInListInt(arr, 88, false)) // 88: [10 9]
	log.Println("88:", ExistsInListInt(arr, 88, true))  // 88: [10]
	log.Println("89:", ExistsInListInt(arr, 89, false)) // 89: [-1]
}
