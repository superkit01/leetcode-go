package dweek130

func satisfiesConditions(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	if m == 1 && n == 1 {
		return true
	}
	if m == 1 {
		for i := 0; i < n-1; i++ {
			if grid[0][i] == grid[0][i+1] {
				return false
			}
		}
	}

	if n == 1 {
		for i := 0; i < m-1; i++ {
			if grid[i][0] != grid[i+1][0] {
				return false
			}
		}

	}

	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != grid[i+1][j] {
				return false
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n-1; j++ {
			if grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true

}
