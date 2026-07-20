package easy

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])

	k = k % (m * n)
	if k == 0 {
		return grid
	}

	temp := make([]int, m*n*2)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp[i*n+j] = grid[i][j]
		}
	}
	for i := 0; i < m*n; i++ {
		temp[i+m*n] = temp[i]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = temp[m*n-k+i*n+j]
		}
	}

	return grid

}

