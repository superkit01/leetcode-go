package middle

func countServers(grid [][]int) int {
	col := make([]int, len(grid))
	row := make([]int, len(grid[0]))

	for i := range grid {
		for j := range grid[i] {
			col[i] += grid[i][j]
			row[j] += grid[i][j]
		}
	}
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && (col[i] > 1 || row[j] > 1) {
				ans++
			}
		}
	}
	return ans
}
