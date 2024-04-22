package lcr

import "math"

func divide(a int, b int) int {

	if b == 0 {
		return math.MaxInt32
	}
	if b == 1 {
		return a
	}

	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	if b == -1 {
		return -a
	}

	if a < 0 {
		return -divide(-a, b)
	}
	if b < 0 {
		return -divide(a, -b)
	}

	if a < b {
		return 0
	}

	res := 1

	temp := b
	for a-temp > temp {
		res = res << 1
		temp = temp << 1
	}
	return res + divide(a-temp, b)

}
