package lcr

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int, grid [][]int) int

	dfs = func(i, j int, grid [][]int) int {
		cnt := 1
		grid[i][j] = 0
		if i-1 >= 0 && grid[i-1][j] == 1 {
			cnt += dfs(i-1, j, grid)
		}
		if i+1 < m && grid[i+1][j] == 1 {
			cnt += dfs(i+1, j, grid)
		}
		if j-1 >= 0 && grid[i][j-1] == 1 {
			cnt += dfs(i, j-1, grid)
		}
		if j+1 < n && grid[i][j+1] == 1 {
			cnt += dfs(i, j+1, grid)
		}
		return cnt
	}

	maxCnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cnt := dfs(i, j, grid)
				if maxCnt < cnt {
					maxCnt = cnt
				}
			}
		}
	}
	return maxCnt

}
