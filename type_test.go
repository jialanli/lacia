package lacia

import (
	"log"
	"testing"
)

func TestGetValTypeOf(t *testing.T) {
	log.Println(GetValTypeOf(123))
	log.Println(GetValTypeOf(true))
	log.Println(GetValTypeOf(1))
	log.Println(GetValTypeOf([]int{1, 2, 3}))
	log.Println(GetValTypeOf([]int64{1, 2, 3}))
	log.Println(GetValTypeOf([]string{"1", "A", "C"}))
	log.Println(GetValTypeOf([]float64{1.32, 0.9, -0.05}))
	log.Println(GetValTypeOf([]float32{0.333333, 0.9, -0.05}))
	log.Println(GetValTypeOf(0.333333))
	log.Println(GetValTypeOf(func() {}))            // func()
	log.Println(GetValTypeOf(func(int, string) {})) // func(int, string)
	log.Println(GetValTypeOf(func(string) {}))      // func(string)
	log.Println(GetValTypeOf(FuncTest))             // func(string, int) ([]int, concia)

	ch1 := make(chan int)
	var ch2 chan string
	var ch3 chan StructTest
	var ch4 chan []StructTest
	log.Println(GetValTypeOf(ch1)) // chan int
	log.Println(GetValTypeOf(ch2)) // chan string
	log.Println(GetValTypeOf(ch3)) // chan lacia.StructTest
	log.Println(GetValTypeOf(ch4)) // chan []lacia.StructTest

	m1 := make(map[string]int)
	m2 := make(map[string]StructTest)
	m3 := make(map[string][]StructTest)
	m4 := make(map[int]map[string]int)
	m5 := make(map[int]func(int, string) (int, int))
	var m6 map[int]struct{}
	log.Println(GetValTypeOf(m1)) // map[string]int
	log.Println(GetValTypeOf(m2)) // map[string]lacia.StructTest
	log.Println(GetValTypeOf(m3)) // map[string][]lacia.StructTest
	log.Println(GetValTypeOf(m4)) // map[int]map[string]int
	log.Println(GetValTypeOf(m5)) // map[int]func(int, string) (int, int)
	log.Println(GetValTypeOf(m6)) // map[int]struct {}
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
	log.Println(res)
}
