package lcr

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}

	queue := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				ans[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	value := 1
	for len(queue) > 0 {
		tempQueue := [][2]int{}
		for _, v := range queue {
			if v[0]-1 >= 0 && ans[v[0]-1][v[1]] == -1 {
				ans[v[0]-1][v[1]] = value
				tempQueue = append(tempQueue, [2]int{v[0] - 1, v[1]})
			}

			if v[0]+1 < m && ans[v[0]+1][v[1]] == -1 {
				ans[v[0]+1][v[1]] = value
				tempQueue = append(tempQueue, [2]int{v[0] + 1, v[1]})
			}

			if v[1]-1 >= 0 && ans[v[0]][v[1]-1] == -1 {
				ans[v[0]][v[1]-1] = value
				tempQueue = append(tempQueue, [2]int{v[0], v[1] - 1})
			}

			if v[1]+1 < n && ans[v[0]][v[1]+1] == -1 {
				ans[v[0]][v[1]+1] = value
				tempQueue = append(tempQueue, [2]int{v[0], v[1] + 1})
			}
		}
		value++
		queue = tempQueue
	}

	return ans

}
