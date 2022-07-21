package main

func shiftGrid(grid [][]int, k int) [][]int {

	m := len(grid)
	n := len(grid[0])

	temp := make([]int, 0)
	for i := 0; i < m; i++ {
		temp = append(temp, grid[i]...)
	}
	k = k % len(temp)

	newTemp := make([]int, len(temp))
	for i, v := range temp {
		index := i + k
		if index > len(temp)-1 {
			index = index - len(temp)
		}
		newTemp[index] = v
	}

	result := make([][]int, 0)
	for i := 0; i < m; i++ {
		result = append(result, make([]int, n))
	}

	for i, v := range newTemp {
		r := i / n
		c := i - r*n
		result[r][c] = v
	}

	return result

}

// func main() {
// shiftGrid([][]int{{1}}, 100)
// }
