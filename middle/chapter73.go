package middle

func setZeroes(matrix [][]int) {
	col := make([]bool, len(matrix))
	row := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				col[i] = true
				row[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if col[i] || row[j] {
				matrix[i][j] = 0
			}
		}
	}

}
