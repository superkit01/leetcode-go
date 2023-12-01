package middle

func permuteUnique(nums []int) [][]int {

	result := make([][]int, 0)

	recursive1(make([]int, 0), nums, &result)

	return result
}

func recursive1(pre []int, remaining []int, result *[][]int) {

	if len(remaining) == 0 {
		*result = append(*result, pre)
		return
	}

	cache := make(map[int]bool, 0)
	for i := 0; i < len(remaining); i++ {
		if _, ok := cache[remaining[i]]; ok {
			continue
		}
		cache[remaining[i]] = true

		newRemaining := make([]int, 0)
		if i == 0 {
			newRemaining = append(newRemaining, remaining[1:]...)
		} else if i == len(remaining)-1 {
			newRemaining = append(newRemaining, remaining[:len(remaining)-1]...)
		} else {
			newRemaining = append(newRemaining, remaining[0:i]...)
			newRemaining = append(newRemaining, remaining[i+1:]...)
		}

		newPre := make([]int, 0)
		newPre = append(newPre, pre...)
		newPre = append(newPre, remaining[i])
		recursive1(newPre, newRemaining, result)

	}

}
