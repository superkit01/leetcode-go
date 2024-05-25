package dweek131

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	pos := []int{}

	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}

	ans := make([]int, len(queries))

	for i, v := range queries {
		if v > len(pos) {
			ans[i] = -1
		} else {
			ans[i] = pos[v-1]
		}
	}
	return ans

}
