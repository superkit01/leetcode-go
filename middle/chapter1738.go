package middle

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	xor := make([][]int, m+1)
	for i := range xor {
		xor[i] = make([]int, n+1)
	}

	ans := make([]int, m*n)
	index := 0

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			xor[i][j] = xor[i-1][j] ^ xor[i][j-1] ^ xor[i-1][j-1] ^ matrix[i-1][j-1]

			ans[index] = xor[i][j]
			index++
		}
	}

	sort.Ints(ans)

	return ans[len(ans)-k]

}
