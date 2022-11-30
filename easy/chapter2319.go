package easy

func checkXMatrix(grid [][]int) bool {
	m := len(grid)
	for i, v := range grid {
		for j, w := range v {

			if (i == j || i+j == m-1) && w == 0 {
				return false
			}

			if (i != j && i+j != m-1) && w != 0 {
				return false
			}
		}
	}
	return true
}
