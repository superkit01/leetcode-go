package topics

import "math"

//快速幂
//  x * x * x * x * x  ~ n
func quickX(x, n int) int {
	// n  1000101101
	if n == 0 {
		return 1
	}

	ans := 1
	if n&0b1 == 0b1 {
		ans = quickX(x, n-1) * x
	} else {
		ans = quickX(x, n>>1) * quickX(x, n>>1)
	}
	return ans
}

//快速乘
//  x + x + x + x + x ~ n
func quickA(x, n int) int {
	if n == 0 {
		return 0
	}
	ans := 0
	if n&0b1 == 0b1 {
		ans = quickA(x, n-1) + x
	} else {
		ans = quickA(x, n>>1) + quickA(x, n>>1)
	}
	return ans
}

//快速除 (递归)
func quickD(x, n int) int {
	if n == 0 {
		return math.MaxInt32
	}
	if n == 1 {
		return x
	}

	if x < n {
		return 0
	}

	res := 1
	tmp := n
	for x-tmp > tmp {
		res <<= 1
		tmp <<= 1
	}

	return res + quickD(x-tmp, n)

}

//快速除 (迭代)
func quickD2(x, n int) int {
	if n == 0 {
		return math.MaxInt32
	}
	if n == 1 {
		return x
	}
	if x < n {
		return 0
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
