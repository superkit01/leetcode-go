package lcr

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	ans := make([]int, n+1)
	ans[0] = 0
	ans[1] = 1
	for i := 2; i <= n; i++ {
		ans[i] = (ans[i-2] + ans[i-1]) % 1000000007
	}
	return ans[n]

}
