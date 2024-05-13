package middle

//BFS

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := make([][]int, 0)
	cnt1 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
			if grid[i][j] == 1 {
				cnt1++
			}
		}
	}
	ans := 0
	for len(queue) > 0 {
		tempQueue := make([][]int, 0)
		for _, v := range queue {
			if v[0]-1 >= 0 && grid[v[0]-1][v[1]] == 1 {
				tempQueue = append(tempQueue, []int{v[0] - 1, v[1]})
				grid[v[0]-1][v[1]] = 2
				cnt1--
			}
			if v[0]+1 < m && grid[v[0]+1][v[1]] == 1 {
				tempQueue = append(tempQueue, []int{v[0] + 1, v[1]})
				grid[v[0]+1][v[1]] = 2
				cnt1--
			}

			if v[1]-1 >= 0 && grid[v[0]][v[1]-1] == 1 {
				tempQueue = append(tempQueue, []int{v[0], v[1] - 1})
				grid[v[0]][v[1]-1] = 2
				cnt1--
			}

			if v[1]+1 < n && grid[v[0]][v[1]+1] == 1 {
				tempQueue = append(tempQueue, []int{v[0], v[1] + 1})
				grid[v[0]][v[1]+1] = 2
				cnt1--
			}
		}
		if len(tempQueue) > 0 {
			ans++
		}
		queue = tempQueue
	}

	if cnt1 == 0 {
		return ans
	} else {
		return -1
	}

}
