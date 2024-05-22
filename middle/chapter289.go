package middle

func gameOfLife(board [][]int) {
	direction := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	cache := make(map[[2]int]int)
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveCnt := 0
			deadCnt := 0
			for _, v := range direction {
				if i+v[0] < 0 || i+v[0] >= m || j+v[1] < 0 || j+v[1] >= n {
					continue
				}
				if ori, ok := cache[[2]int{i + v[0], j + v[1]}]; ok {
					if ori == 0 {
						deadCnt++
					} else {
						liveCnt++
					}
				} else {
					if board[i+v[0]][j+v[1]] == 0 {
						deadCnt++
					} else {
						liveCnt++
					}
				}

			}

			if board[i][j] == 1 && (liveCnt < 2 || liveCnt > 3) {
				cache[[2]int{i, j}] = 1
				board[i][j] = 0
			}

			if board[i][j] == 0 && liveCnt == 3 {
				cache[[2]int{i, j}] = 0
				board[i][j] = 1
			}
		}
	}

}
