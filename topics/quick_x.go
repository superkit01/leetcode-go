package topics

//快速幂
//  x * x * x * x * x  ~ n
func quickX(x, n int) int {
	// n  1000101101
	if n == 0 {
		return 1
	}

	ans := 1
	if n%2 == 1 {
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
	if n%2 == 1 {
		ans = quickA(x, n-1) + x
	} else {
		ans = quickA(x, n/2) + quickA(x, n/2)
	}
	return ans
}
