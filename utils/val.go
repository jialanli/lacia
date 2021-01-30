package lacia

func SwapValInt64(x, y *int64) {
	*x, *y = *y, *x
}

func SwapValInt(x, y *int) {
	*x, *y = *y, *x
}

func SwapValString(x, y *string) {
	*x, *y = *y, *x
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// rate = a-b/b
func CalcRateInt64(src, dst int64) float64 {
	if dst == 0 {
		return -1
	}
	return float64(src-dst) / float64(dst)
}

func CalcRateFloat64(src, dst float64) float64 {
	if int(dst) == 0 {
		return -1
	}
	return (src - dst) / dst
}

// 整数反转
func ReverseInt(x int) int {
	num, res := x, 0
	for num != 0 {
		res, num = res*10+(num%10), num/10
	}

	if res < -2147483648 || res > 2147483647 {
		return 0
	}

	return res
}
