package middle

import "math"

func divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}
	if dividend > 0 {
		return -divide(-dividend, divisor)
	}
	if divisor > 0 {
		return -divide(dividend, -divisor)
	}
	if dividend > divisor {
		return 0
	}
	res := 1
	tmp := divisor
	for (dividend - divisor) <= divisor {
		res = res << 1
		divisor += divisor
	}
	return res + divide(dividend-divisor, tmp)
}
