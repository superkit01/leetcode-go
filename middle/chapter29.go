package middle

import "math"

func Divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}
	if dividend > 0 {
		return -Divide(-dividend, divisor)
	}
	if divisor > 0 {
		return -Divide(dividend, -divisor)
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

	return res + Divide(dividend-divisor, tmp)
}

func Divide2(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}
	if dividend < 0 {
		return -Divide2(-dividend, divisor)
	}
	if divisor < 0 {
		return -Divide2(dividend, -divisor)
	}
	if dividend < divisor {
		return 0
	}
	return quickD(dividend, divisor)
}

func quickD(x, n int) int {
	if n == 0 {
		return math.MaxInt32
	}
	if n == 1 {
		return x
	}

	// [ n, n+n, n+n+n+n, n+n+n+n+n+n+n+n, ....]
	// [ 1, 2,   4,       8,               ....]

	dp := make([]int, 0)
	dp = append(dp, n)
	res := make([]int, 0)
	res = append(res, 1)
	for x-dp[len(dp)-1] > dp[len(dp)-1] {
		dp = append(dp, dp[len(dp)-1]<<1)
		res = append(res, res[len(res)-1]<<1)
	}

	ans := 0
outer:
	for {
		for i := len(dp) - 1; i >= 0; i-- {
			if x >= dp[i] {
				ans += res[i]
				x = x - dp[i]
				continue outer
			}
		}
		return ans
	}

}
