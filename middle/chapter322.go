package middle

import "math"

// 				完全背包
//
// 					 1		2		5	coins i
// j	amount  0	 0		0		0
// 				1	 1
// 				2	 2
// 				3
// 				4
// 				5
// 				...
//
//
//
//
//

func coinChange(coins []int, amount int) int {
	// i对应coin  ; j对应amount
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

	if dp[len(coins)-1][amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[len(coins)-1][amount]
	}

}

// 0 1 2 3 4 5 6 7 8 9 10 11 12
// 0 1 1     1

func coinChangeII(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		minValue := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				minValue = min(minValue, dp[i-coins[j]])
			}
		}
		if minValue == math.MaxInt32 {
			dp[i] = math.MaxInt32
		} else {
			dp[i] = minValue + 1
		}
	}

	if dp[len(dp)-1] == math.MaxInt32 {
		return -1
	} else {
		return dp[len(dp)-1]
	}
}

