package middle

func firstCompleteIndex(arr []int, mat [][]int) int {
	row := make([]int, len(mat))
	col := make([]int, len(mat[0]))

	cache := make(map[int][]int, 0)

	for i, vi := range mat {
		for j, vj := range vi {
			cache[vj] = make([]int, 2)
			cache[vj][0] = i
			cache[vj][1] = j
		}
	}

	for i := 0; i < len(arr); i++ {
		pos := cache[arr[i]]
		row[pos[0]]++
		col[pos[1]]++

		if row[pos[0]] == len(mat[0]) {
			return i
		}

		if col[pos[1]] == len(mat) {
			return i
		}

	}
	return 0

}
