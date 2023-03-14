package middle

import "math"

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	result := make([][]int, len(rowSum))
	for i, _ := range result {
		result[i] = make([]int, len(colSum))
	}
outer:
	for i := 0; i < len(rowSum); i++ {
		for j := 0; j < len(colSum); j++ {
			if rowSum[i] == 0 {
				continue outer
			}
			if colSum[j] == 0 {
				continue
			}
			result[i][j] = int(math.Min(float64(rowSum[i]), float64(colSum[j])))
			rowSum[i] -= result[i][j]
			colSum[j] -= result[i][j]
		}
	}

	return result

}
