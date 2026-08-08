package lcr

func maxSales(sales []int) int {
	dp := make([]int, len(sales))

	dp[0] = sales[0]

	for i := 1; i < len(sales); i++ {
		dp[i] = max(dp[i-1]+sales[i], sales[i])
	}

	ans := dp[0]
	for i := 1; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}

	return ans

}
