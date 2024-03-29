package middle

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}
	if m == 1 {
		for i := 0; i < n; i++ {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}
		return 1
	}

	if n == 1 {
		for i := 0; i < m; i++ {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
		return 1
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 || dp[0][i-1] == 0 {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				if dp[i-1][j] == 0 || obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]

}

// func main() {
// 	uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
// }
