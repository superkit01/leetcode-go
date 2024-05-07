package lcs

func largestArea(grid []string) int {
	m := len(grid)
	n := len(grid[0])

	mark := make([][]int, m)
	for i := range mark {
		mark[i] = make([]int, n)
	}

	//标记与走廊直接相邻的主题
	for i, v := range grid {
		for j := 0; j < len(v); j++ {
			if grid[i][j] == '0' {
				mark[i][j] = -1
				if i-1 >= 0 {
					dfs(grid, &mark, i-1, j, m, n)
				}
				if i+1 < m {
					dfs(grid, &mark, i+1, j, m, n)
				}
				if j-1 >= 0 {
					dfs(grid, &mark, i, j-1, m, n)
				}
				if j+1 < n {
					dfs(grid, &mark, i, j+1, m, n)
				}
				continue
			}

			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				dfs(grid, &mark, i, j, m, n)
			}
		}
	}

	//去除标记之后计算主题的面积
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mark[i][j] == -1 {
				continue
			}
			ans = max(ans, dfs2(grid, &mark, i, j, m, n))
		}
	}

	return ans

}

// dfs 标记直接相邻区域
func dfs(grid []string, mark *[][]int, i, j, m, n int) {
	if (*mark)[i][j] == -1 {
		return
	}
	(*mark)[i][j] = -1
	if i-1 >= 0 && grid[i][j] == grid[i-1][j] {
		dfs(grid, mark, i-1, j, m, n)
	}
	if i+1 < m && grid[i][j] == grid[i+1][j] {
		dfs(grid, mark, i+1, j, m, n)
	}
	if j-1 >= 0 && grid[i][j] == grid[i][j-1] {
		dfs(grid, mark, i, j-1, m, n)
	}
	if j+1 < n && grid[i][j] == grid[i][j+1] {
		dfs(grid, mark, i, j+1, m, n)
	}
}

// dfs 计算数量
func dfs2(grid []string, mark *[][]int, i, j, m, n int) int {
	if (*mark)[i][j] == -1 {
		return 0
	}
	(*mark)[i][j] = -1

	ans := 1
	if i-1 >= 0 && grid[i][j] == grid[i-1][j] {
		ans += dfs2(grid, mark, i-1, j, m, n)
	}
	if i+1 < m && grid[i][j] == grid[i+1][j] {
		ans += dfs2(grid, mark, i+1, j, m, n)
	}
	if j-1 >= 0 && grid[i][j] == grid[i][j-1] {
		ans += dfs2(grid, mark, i, j-1, m, n)
	}
	if j+1 < n && grid[i][j] == grid[i][j+1] {
		ans += dfs2(grid, mark, i, j+1, m, n)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
