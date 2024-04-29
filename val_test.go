package lacia

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestReverseInt(t *testing.T) {
	log.Println(ReverseInt(123))  // 321
	log.Println(ReverseInt(-123)) // -321
	log.Println(ReverseInt(210))  // 12
}

func TestSwapValString(t *testing.T) {
	a, b := 1, 3
	SwapValInt(&a, &b)
	log.Println(a, b)
	c, d := "C", "D"
	SwapValString(&c, &d)
	log.Println(c, d)

	var e, f int64 = 11, 3
	SwapValInt64(&e, &f)
	log.Println(e, f)
}

type R int

func TestSw(t *testing.T) {
	a := 5
	b := R(2)
	log.Println(GetValTypeOf(a))
	log.Println(GetValTypeOf(b))
	c := a + int(b)
	log.Println(c)

}

func TestA(t *testing.T) {
	type Obj struct {
		A string
		B string
	}
	var m sync.Map

	go func() {
		m.Store(Obj{A: "a", B: "1"}, "x")
		m.Store(Obj{A: "b", B: "2"}, "")
		m.Store(Obj{A: "d", B: "5"}, "1")
	}()
	go func() {
		m.Store(Obj{A: "a", B: "2"}, "y")
		m.Store(Obj{A: "c", B: "3"}, "")
		m.Store(Obj{A: "d", B: "5"}, "2")
	}()

	time.Sleep(time.Second * 3)
	m.Range(func(key, value interface{}) bool {
		log.Println(key, value)
		return true
	})
}
