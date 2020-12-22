package lacia

import "fmt"

type ArrList struct {
	ID  int64
	Num int // 下标
}

// 数组去重算法：一次遍历拿到重复的元素、重复元素的个数&列表&下标、不重复元素个数&列表
func GetDataNoRepeat(list []int64) (err error) {
	var noRepeatList []int64
	var repeatList []int64
	m := make(map[int64][]ArrList)
	for i, ID := range list {
		obj := ArrList{ID: ID, Num: i}
		if data, ok := m[ID]; ok {
			l := append(data, obj)
			m[ID] = l
			repeatList = append(repeatList, ID)
			continue
		}
		m[ID] = []ArrList{obj}
		noRepeatList = append(noRepeatList, ID)
	}

	fmt.Println("重复的ID个数：", len(repeatList), "; 重复的有：", repeatList) // 10
	fmt.Println("不重复的个数：", len(noRepeatList))                       // 5493
	return
}

// 给定一个列表,去除/找出重复元素
func RepeatElement(list interface{}) (repeatList interface{}, repeatCount int) {
	//var inputList []interface{}

	return
}
