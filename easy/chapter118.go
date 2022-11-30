package easy

func generate(numRows int) [][]int {
	result := make([][]int, 0)

	current := []int{1}
	result = append(result, current)
	if numRows == 1 {
		return result
	}

	current = []int{1, 1}
	result = append(result, current)
	if numRows == 2 {
		return result
	}

	for i := 3; i <= numRows; i++ {
		tmp := result[len(result)-1]

		temp2 := make([]int, 0)
		for j := 0; j < i-1; j++ {
			if j == 0 {
				temp2 = append(temp2, 1)
			} else {
				temp2 = append(temp2, tmp[j-1]+tmp[j])
			}
		}
		temp2 = append(temp2, 1)
		result = append(result, temp2)

	}
	return result
}
