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
