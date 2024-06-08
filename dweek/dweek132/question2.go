package dweek132

func findWinningPlayer(skills []int, k int) int {
	max := skills[0]
	index := 0
	count := 0
	for i := 1; i < len(skills); i++ {
		if count == k {
			break
		}
		if max > skills[i] {
			count++
		} else {
			max = skills[i]
			index = i
			count = 1
		}
	}
	return index
}
