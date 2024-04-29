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

// growth rate
func CalcRateInt64(thisCycle, previousCycle int64) float64 {
	if previousCycle == 0 {
		return -1
	}
	return float64(thisCycle-previousCycle) / float64(previousCycle)
}

func CalcRateFloat64(thisCycle, previousCycle float64) float64 {
	if int(previousCycle) == 0 {
		return -1
	}
	return (thisCycle - previousCycle) / previousCycle
}

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
