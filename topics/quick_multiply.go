package topics

//快速幂
func quickMultiply(x, n int) int {
	// n  1000101101

	if n == 0 {
		return 1
	}

	ans := 1
	if n%2 == 1 {
		ans = quickMultiply(x, n-1) * x
	} else {
		ans = quickMultiply(x, n>>1) * quickMultiply(x, n>>1)
	}
	return ans

}
