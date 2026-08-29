package lcr

func SpiralArray(array [][]int) []int {
	if len(array) == 0 || len(array[0]) == 0 {
		return []int{}
	}
	m, n := len(array), len(array[0])

	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}

	ans := make([]int, m*n)

	x, y := 0, 0
	index := 0
	for i := 0; i < m*n; i++ {

		ans[i] = array[x][y]
		visited[x][y] = 1

		nextX := x + direction[index][1]
		nextY := y + direction[index][0]

		if nextX >= m || nextY >= n || nextX < 0 || nextY < 0 || visited[nextX][nextY] == 1 {
			index = (index + 1) % 4
		}

		x += direction[index][1]
		y += direction[index][0]

	}
	return ans

}
