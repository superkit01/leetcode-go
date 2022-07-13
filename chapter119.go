package main

func getRow(rowIndex int) []int {
	temp := []int{1}

	if rowIndex == 0 {
		return temp
	}

	for i := 1; i <= rowIndex; i++ {
		current := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				current = append(current, 1)
			} else {
				current = append(current, temp[j-1]+temp[j])
			}
		}
		temp = current
	}

	return temp
}
