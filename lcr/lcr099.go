package lcr

func minPathSum(grid [][]int) int {
	if len(grid[0]) > 0 {
		for i := 1; i < len(grid[0]); i++ {
			grid[0][i] = grid[0][i] + grid[0][i-1]
		}
	}
	if len(grid) > 0 {
		for i := 1; i < len(grid); i++ {
			grid[i][0] = grid[i][0] + grid[i-1][0]
		}
	}

	if len(grid) > 1 && len(grid[0]) > 1 {
		for i := 1; i < len(grid); i++ {
			for j := 1; j < len(grid[0]); j++ {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]

}
