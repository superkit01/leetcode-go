package middle

// 完全背包
//    0 1 2 3 4 5 ...
//  1 1 1
//  2 1
//  5 1
func change(amount int, coins []int) int {

	dp := make([][]int, len(coins))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			//不使用当前coin
			if i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			}
			//使用当前coin
			if j-coins[i] >= 0 {
				dp[i][j] += dp[i][j-coins[i]]
			}
		}
	}
	return dp[len(dp)-1][amount]
}
