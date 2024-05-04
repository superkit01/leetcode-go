package lcr

func wardrobeFinishing(m int, n int, cnt int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if digit(i) > cnt || dp[0][i-1] == 0 {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		if digit(i) > cnt || dp[i-1][0] == 0 {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] == 0 && dp[i][j-1] == 0 {
				dp[i][j] = 0
			} else if digit(i)+digit(j) > cnt {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans += dp[i][j]
		}
	}

	return ans
}

func digit(x int) int {
	ans := 0
	for x > 0 {
		ans += x % 10
		x /= 10
	}
	return ans
}
