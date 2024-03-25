package middle

import "math"

//   0 1 2 3 4 5 ... amount
// 1 0
// 2 0
// 5 0

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 0
	}

	for j := 1; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			count := math.MaxInt32
			//不使用当前coin

			if i-1 >= 0 {
				count = int(math.Min(float64(count), float64(dp[i-1][j])))
			}
			//使用当前coin
			if j-coins[i] >= 0 {
				count = int(math.Min(float64(count), float64(dp[i][j-coins[i]]+1)))

			}
			dp[i][j] = count
		}
	}

	if dp[len(coins)-1][amount] ==math.MaxInt32{
		return -1
	}else{
		return dp[len(coins)-1][amount]
	}

}
