package week401

func valueAfterKSeconds(n int, k int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i <= k; i++ {
		temp := make([]int, n)
		temp[0] = 1
		for j := 1; j < n; j++ {
			temp[j] = (temp[j-1] + dp[j]) % 1000000007
		}
		dp = temp
	}
	return dp[n-1]
}
