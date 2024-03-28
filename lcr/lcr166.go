package lcr

import "math"

//	二维数组 最大路径 之和
//  dp[i][j] = max( dp[i-1][j], dp[i][j-1]) + grid[i][j]

func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j != 0 {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}

			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]

}
