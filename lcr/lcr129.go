package lcr

func wordPuzzle(grid [][]byte, target string) bool {
	m := len(grid)
	n := len(grid[0])

	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == target[index] {
				matrix[i][j] = 1
				if dfs129(i, j, index, matrix, grid, target) {
					return true
				}
				matrix[i][j] = 0
			}
		}
	}
	return false

}

func dfs129(i, j, index int, matrix [][]int, grid [][]byte, target string) bool {
	if index == len(target)-1 {
		return true
	}

	if i+1 < len(grid) && grid[i+1][j] == target[index+1] && matrix[i+1][j] == 0 {
		matrix[i+1][j] = 1
		if dfs129(i+1, j, index+1, matrix, grid, target) {
			return true
		}
		matrix[i+1][j] = 0
	}

	if i-1 >= 0 && grid[i-1][j] == target[index+1] && matrix[i-1][j] == 0 {
		matrix[i-1][j] = 1
		if dfs129(i-1, j, index+1, matrix, grid, target) {
			return true
		}
		matrix[i-1][j] = 0
	}

	if j+1 < len(grid[0]) && grid[i][j+1] == target[index+1] && matrix[i][j+1] == 0 {
		matrix[i][j+1] = 1
		if dfs129(i, j+1, index+1, matrix, grid, target) {
			return true
		}
		matrix[i][j+1] = 0
	}

	if j-1 >= 0 && grid[i][j-1] == target[index+1] && matrix[i][j-1] == 0 {
		matrix[i][j-1] = 1
		if dfs129(i, j-1, index+1, matrix, grid, target) {
			return true
		}
		matrix[i][j-1] = 0
	}

	return false
}
