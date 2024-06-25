package hard

func goodSubsetofBinaryMatrix(grid [][]int) []int {

	m, n := len(grid), len(grid[0])
	arr := make([]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[i] |= grid[i][j] << j
		}
		if arr[i] == 0 {
			return []int{i}
		}
	}

	i := 1
	for i < m {
		for j := 0; j < i; j++ {
			if arr[i]&arr[j] == 0 {
				return []int{j, i}
			}
		}
		i++
	}
	return []int{}

}
