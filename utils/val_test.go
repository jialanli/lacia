package lacia

import (
	"fmt"
	"testing"
)

func TestReverseInt(t *testing.T) {
	fmt.Println(ReverseInt(123))  // 321
	fmt.Println(ReverseInt(-123)) // -321
	fmt.Println(ReverseInt(210))  // 12
}

func TestSwapValString(t *testing.T) {
	a, b := 1, 3
	SwapValInt(&a, &b)
	fmt.Println(a, b)
	c, d := "C", "D"
	SwapValString(&c, &d)
	fmt.Println(c, d)

	var e, f int64 = 11, 3
	SwapValInt64(&e, &f)
	fmt.Println(e, f)
}
