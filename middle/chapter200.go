package middle

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				grid[i][j] = '0'

				//BFS  广度优先搜索

				queue := make([][]int, 0)
				queue = append(queue, []int{i, j})
				for len(queue) > 0 {
					temp := make([][]int, 0)
					for _, v := range queue {
						if v[0]-1 >= 0 && grid[v[0]-1][v[1]] == '1' {
							grid[v[0]-1][v[1]] = '0'
							temp = append(temp, []int{v[0] - 1, v[1]})
						}
						if v[0]+1 < m && grid[v[0]+1][v[1]] == '1' {
							grid[v[0]+1][v[1]] = '0'
							temp = append(temp, []int{v[0] + 1, v[1]})
						}
						if v[1]-1 >= 0 && grid[v[0]][v[1]-1] == '1' {
							grid[v[0]][v[1]-1] = '0'
							temp = append(temp, []int{v[0], v[1] - 1})
						}
						if v[1]+1 < n && grid[v[0]][v[1]+1] == '1' {
							grid[v[0]][v[1]+1] = '0'
							temp = append(temp, []int{v[0], v[1] + 1})
						}
					}
					queue = temp
				}

			}

		}
	}

	return ans

}

func numIslands2(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	//DFS 深度优先搜索
	var dfs9 func(x, y int)
	dfs9 = func(i, j int) {
		if i < 0 || i >= m {
			return
		}
		if j < 0 || j >= n {
			return
		}

		if grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs9(i-1, j)
			dfs9(i+1, j)
			dfs9(i, j-1)
			dfs9(i, j+1)
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs9(i, j)
			}

		}
	}

	return ans

}
