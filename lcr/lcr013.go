package lcr

type NumMatrix struct {
	matrix [][]int
}

func Constructor013(matrix [][]int) NumMatrix {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j > 0 {
				matrix[i][j] = matrix[i][j-1] + matrix[i][j]
			}
		}
	}
	return NumMatrix{
		matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			ans += this.matrix[i][col2]
		} else {
			ans += this.matrix[i][col2] - this.matrix[i][col1-1]
		}

	}
	return ans
}
