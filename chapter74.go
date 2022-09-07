package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	i := 0
	j := n - 1

	for i >= 0 && i < m && j >= 0 && j < n {

		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}

	}
	return false

}
