package dweek129

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			countB := 0
			if grid[i][j] == 'B' {
				countB++
			}
			if grid[i+1][j] == 'B' {
				countB++
			}
			if grid[i][j+1] == 'B' {
				countB++
			}
			if grid[i+1][j+1] == 'B' {
				countB++
			}
			if countB >= 3 || countB <= 1 {
				return true
			}
		}
	}
	return false
}
