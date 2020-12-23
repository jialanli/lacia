package lacia

import (
	"fmt"
	"testing"
)

func TestSortArr(t *testing.T) {
	var sortArray = []int{11, 3, 6, 25, 15, 1, 27}
	fmt.Println("before sort:", sortArray)
	SortArr(sortArray, "AsC")
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
	var arr = []string{"F","V","R","V","A","P","F","A","F"}
	fmt.Println("before:", arr)
	arr0, m := RepeatElementSidesString(arr)
	fmt.Println("arr0:", arr0)
	fmt.Println("m:", m)
}
