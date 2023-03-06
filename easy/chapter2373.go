package easy

func largestLocal(grid [][]int) [][]int {
	result := make([][]int, len(grid)-2)
	for i := range result {
		result[i] = make([]int, len(grid)-2)
	}

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid)-2; j++ {
			max := grid[i][j]
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					if max < grid[i+m][j+n] {
						max = grid[i+m][j+n]
					}
				}
			}
			result[i][j] = max
		}
	}
	return result
}
