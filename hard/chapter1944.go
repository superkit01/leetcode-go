package hard

func canSeePersonsCount(heights []int) []int {
	monotonicStack := make([]int, 0)
	result := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for len(monotonicStack) >0 && heights[i] > monotonicStack[len(monotonicStack)-1] {
				monotonicStack = monotonicStack[0 : len(monotonicStack)-1]
				result[i]++
		}
		if len(monotonicStack) >0 {
			result[i]++
		}
		monotonicStack=append(monotonicStack, heights[i])
	}
	return result
}
