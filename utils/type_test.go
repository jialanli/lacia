package lacia

import (
	"fmt"
	"testing"
)

func TestGetValTypeOf(t *testing.T) {
	fmt.Println(GetValTypeOf(123))
	fmt.Println(GetValTypeOf(true))
	fmt.Println(GetValTypeOf(1))
	fmt.Println(GetValTypeOf([]int{1, 2, 3}))
	fmt.Println(GetValTypeOf([]int64{1, 2, 3}))
	fmt.Println(GetValTypeOf([]string{"1", "A", "C"}))
	fmt.Println(GetValTypeOf([]float64{1.32, 0.9, -0.05}))
	fmt.Println(GetValTypeOf([]float32{0.333333, 0.9, -0.05}))
	fmt.Println(GetValTypeOf(0.333333))
	fmt.Println(GetValTypeOf(func() {}))            // func()
	fmt.Println(GetValTypeOf(func(int, string) {})) // func(int, string)
	fmt.Println(GetValTypeOf(func(string) {}))      // func(string)
	fmt.Println(GetValTypeOf(FuncTest))             // func(string, int) ([]int, concia)

	ch1 := make(chan int)
	var ch2 chan string
	var ch3 chan StructTest
	var ch4 chan []StructTest
	fmt.Println(GetValTypeOf(ch1)) // chan int
	fmt.Println(GetValTypeOf(ch2)) // chan string
	fmt.Println(GetValTypeOf(ch3)) // chan lacia.StructTest
	fmt.Println(GetValTypeOf(ch4)) // chan []lacia.StructTest

	m1 := make(map[string]int)
	m2 := make(map[string]StructTest)
	m3 := make(map[string][]StructTest)
	m4 := make(map[int]map[string]int)
	m5 := make(map[int]func(int, string) (int, int))
	var m6 map[int]struct{}
	fmt.Println(GetValTypeOf(m1)) // map[string]int
	fmt.Println(GetValTypeOf(m2)) // map[string]lacia.StructTest
	fmt.Println(GetValTypeOf(m3)) // map[string][]lacia.StructTest
	fmt.Println(GetValTypeOf(m4)) // map[int]map[string]int
	fmt.Println(GetValTypeOf(m5)) // map[int]func(int, string) (int, int)
	fmt.Println(GetValTypeOf(m6)) // map[int]struct {}
}

func FuncTest(string2 string, int2 int) (res []int, err error) {
	return
}

type StructTest struct {
	Name string
}

func TestGetValTypeOfSample(t *testing.T) {
	var flag bool
	res := GetValTypeOfSample(flag)
	fmt.Println(res)
}