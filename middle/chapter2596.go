package middle

func checkValidGrid(grid [][]int) bool {
	currentIndex := grid[0][0]
	if currentIndex != 0 {
		return false
	}
	n := len(grid[0])
	currentPosition := []int{0, 0}

	directions := [][]int{{2, 1}, {2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}
outer:
	for {
		if currentIndex == n*n-1 {
			return true
		}
		for _, v := range directions {
			row := v[0] + currentPosition[0]
			col := v[1] + currentPosition[1]
			if row < 0 || row >= n {
				continue
			}
			if col < 0 || col >= n {
				continue
			}

			if grid[row][col] == currentIndex+1 {
				currentPosition = []int{row, col}
				currentIndex++
				continue outer
			}

		}

		return false

	}

}
