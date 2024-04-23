package lcr

func countBits(n int) []int {

	dp := make([]int, n+1)
	dp[0] = 0

	if n == 0 {
		return dp
	}

	for i := 1; i < n+1; i++ {
		if i&0b1 == 0b1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i/2]
		}
	}
	return dp

}
