package lacia

import (
	"fmt"
	"testing"
)

func TestGetDaysNumByTwoDays(t *testing.T) {
	days0 := GetDaysNumByTwoDays(2020, 6, 30, 8, 1) //跳月，返回33
	fmt.Println("总天数1=", days0)
	days1 := GetDaysNumByTwoDays(2020, 6, 30, 7, 2) //跳月，返回3
	fmt.Println("总天数2=", days1)
	days2 := GetDaysNumByTwoDays(2020, 6, 20, 6, 29) //当月，返回10
	fmt.Println("总天数3=", days2)
	days3 := GetDaysNumByTwoDays(2020, 5, 30, 8, 1) //跳月，返回64
	fmt.Println("总天数4=", days3)
}

func TestCompareTwoStrTime(t *testing.T) {
	fmt.Println("compare 20201001 and 20201001: ", CompareTwoStrTime("20201001", "20201001"))
	fmt.Println("compare 20201002 and 20201001: ", CompareTwoStrTime("20201002", "20201001"))
	fmt.Println("compare 20201102 and 20201001: ", CompareTwoStrTime("20201102", "20201001"))
	fmt.Println("compare 20201102 and 20201201: ", CompareTwoStrTime("20201102", "20201001"))
	fmt.Println("compare 202011 and 202010: ", CompareTwoStrTime("202011", "202010"))
	fmt.Println("compare 202101 and 202010: ", CompareTwoStrTime("202101", "202010"))
	fmt.Println("compare 2021 and 2020: ", CompareTwoStrTime("2021", "2020"))

	fmt.Println("compare 2021-01-01 and 2020-12-01: ", CompareTwoStrTime("2021-01-01", "2020-12-01"))
	fmt.Println("compare 2020-12-01 and 2020-12-01: ", CompareTwoStrTime("2021-01-01", "2020-12-01"))
	fmt.Println("compare 2020-12 and 2020-12: ", CompareTwoStrTime("2020-12", "2020-12"))
	fmt.Println("compare 2020-10-10 18:18:18 and 2020-10-10 18:18:18: ", CompareTwoStrTime("2020-10-10 18:18:18", "2020-10-10 18:18:18"))
	fmt.Println("compare 2020-10-10 18:18:19 and 2020-10-10 18:18:18: ", CompareTwoStrTime("2020-10-10 18:18:19", "2020-10-10 18:18:18"))
	fmt.Println("compare 2020-10-10 18:19:19 and 2020-10-10 18:18:19: ", CompareTwoStrTime("2020-10-10 18:19:19", "2020-10-10 18:18:19"))

	fmt.Println("compare 20201010 18:19:19 and 20201011 18:18:19: ", CompareTwoStrTime("20201010 18:19:19", "20201011 18:18:19"))
	fmt.Println("compare 20201010 18:19:19 and 20201010 18:18:19: ", CompareTwoStrTime("20201010 18:19:19", "20201010 18:18:19"))
	fmt.Println("compare 20201010 18:18:18 and 20201010 18:18:18: ", CompareTwoStrTime("20201010 18:18:18", "20201010 18:18:18"))

	fmt.Println("compare 2020/10/10 18:18:18 and 2020/10/10 18:18:18: ", CompareTwoStrTime("2020/10/10 18:18:18", "2020/10/10 18:18:18"))
	fmt.Println("compare 2020/10/10 18:18:19 and 2020/10/10 18:18:18: ", CompareTwoStrTime("2020/10/10 18:18:19", "2020/10/10 18:18:18"))
	fmt.Println("compare 2020/10/10 and 2020/10/10: ", CompareTwoStrTime("2020/10/10", "2020/10/10"))
	fmt.Println("compare 2020/10/11 and 2020/10/10: ", CompareTwoStrTime("2020/10/11", "2020/10/10"))
	fmt.Println("compare 2020/10 and 2020/10: ", CompareTwoStrTime("2020/10", "2020/10"))
	fmt.Println("compare 2020/10 and 2020/11: ", CompareTwoStrTime("2020/10", "2020/11"))
}
