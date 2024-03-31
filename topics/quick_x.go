package topics

//快速幂
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
