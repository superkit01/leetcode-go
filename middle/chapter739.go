package middle

func dailyTemperatures(temperatures []int) []int {

	//单调栈  [  [temperature, index], ]
	monotonicStack := make([][]int, 0)

	result := make([]int, len(temperatures))

outer:
	for i := len(temperatures) - 1; i >= 0; i-- {
	inner:
		for {
			if len(monotonicStack) == 0 {
				result[i] = 0
				monotonicStack = append(monotonicStack, []int{temperatures[i], i})
				continue outer
			}

			if temperatures[i] >= monotonicStack[len(monotonicStack)-1][0] {

				monotonicStack = monotonicStack[0 : len(monotonicStack)-1]
				continue inner
			} else {
				result[i] = monotonicStack[len(monotonicStack)-1][1] - i
				monotonicStack = append(monotonicStack, []int{temperatures[i], i})
				continue outer
			}
		}

	}

	return result

}
