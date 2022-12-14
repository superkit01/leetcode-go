package middle

func shortestPathBinaryMatrix(grid [][]int) int {
	length := len(grid)
	if grid[0][0] == 1 {
		return -1
	}

	if length == 1 {
		return 1
	}

	value := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	cache := make([][]bool, length)
	for i := range cache {
		cache[i] = make([]bool, length)
	}
	cache[0][0] = true

	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	result := 0

	for len(queue) > 0 {
		result++
		newTemp := make([][]int, 0)
		if cache[length-1][length-1] {
			return result
		}

		for _, v := range queue {
			for _, i := range value {
				if v[0]+i[0] < 0 || v[0]+i[0] >= length {
					continue
				}
				if v[1]+i[1] < 0 || v[1]+i[1] >= length {
					continue
				}

				if grid[v[0]+i[0]][v[1]+i[1]] == 1 {
					continue
				}

				if cache[v[0]+i[0]][v[1]+i[1]] {
					continue
				}

				cache[v[0]+i[0]][v[1]+i[1]] = true
				newTemp = append(newTemp, []int{v[0] + i[0], v[1] + i[1]})
			}
		}
		queue = newTemp
	}

	if cache[length-1][length-1] == true {
		return result
	} else {
		return -1
	}

}
