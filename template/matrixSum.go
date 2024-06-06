package template

type NumMatrix struct {
	matrix [][]int
}

func ConstructorII(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	prefixSum := make([][]int, m+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			//	二维数组前缀和
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + matrix[i-1][j-1]
			//	二维数组前缀异或和
			//prefixSum[i][j] = prefixSum[i-1][j] ^ prefixSum[i][j-1] ^ prefixSum[i-1][j-1] ^ matrix[i-1][j-1]

		}
	}
	return NumMatrix{matrix: prefixSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrix[row2+1][col2+1] - this.matrix[row2+1][col1] - this.matrix[row1][col2+1] + this.matrix[row1][col1]
}
