package dweek129

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	n := len(grid[0])
	if m == 1 || n == 1 {
		return 0
	}

	mSum := make([]int, m)
	nSum := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				mSum[i]++
				nSum[j]++
			}
		}
	}

	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = (mSum[i] - 1) * (nSum[j] - 1)
			}
		}
	}

	var sum int64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum = sum + int64(dp[i][j])
		}
	}

	return sum

}
