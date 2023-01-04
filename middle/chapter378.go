package middle

import "math"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	stack := make([]int, n)
	for i := 0; i < n; i++ {
		stack[i] = 0
	}
	result := 0
	for t := 0; t < k; t++ {
		min := stack[0]
		minV := math.MaxInt
		for i := 0; i < n; i++ {
			if stack[i] >= n {
				continue
			}
			if matrix[i][stack[i]] < minV {
				minV = matrix[i][stack[i]]
				min = i
			}
		}
		stack[min]++
		result = minV
	}
	return result

}
