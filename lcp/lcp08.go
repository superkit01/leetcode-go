package lcp

func getTriggerTime(increase [][]int, requirements [][]int) []int {
	newIncrease := make([][]int, 0)
	newIncrease = append(newIncrease, []int{0, 0, 0})
	newIncrease = append(newIncrease, increase...)

	for i := 1; i < len(newIncrease); i++ {
		newIncrease[i][0] = newIncrease[i][0] + newIncrease[i-1][0]
		newIncrease[i][1] = newIncrease[i][1] + newIncrease[i-1][1]
		newIncrease[i][2] = newIncrease[i][2] + newIncrease[i-1][2]
	}

	result := make([]int, 0)

	for _, v := range requirements {

		i := search2(newIncrease, v)
		result = append(result, i)
	}

	return result

}

func search2(increase [][]int, v []int) int {

	if increase[len(increase)-1][0] < v[0] || increase[len(increase)-1][1] < v[1] || increase[len(increase)-1][2] < v[2] {
		return -1
	}

	l := 0
	r := len(increase) - 1

	for l < r {
		mid := (l + r) / 2

		value := increase[mid]
		if value[0] < v[0] || value[1] < v[1] || value[2] < v[2] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
