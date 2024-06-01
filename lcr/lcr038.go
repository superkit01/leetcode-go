package lcr

import "math"

func DailyTemperatures(temperatures []int) []int {
	stack := [][2]int{{math.MaxInt, 0}}

	ans := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for stack[len(stack)-1][0] <= temperatures[i] {
			stack = stack[0 : len(stack)-1]
		}
		ans[i] = max(0,stack[len(stack)-1][1] -i)
		stack = append(stack, [2]int{temperatures[i], i})
	}
	return ans

}
