package middle

//1,2,1  1 2 1
//0 1 2  3 4 5
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}

	nums = append(nums, nums...)

	monotonicStack := make([]int, 0)

outer:
	for i, v := range nums {
		for {
			if len(monotonicStack) == 0 {
				monotonicStack = append(monotonicStack, i%n)
				continue outer
			}
			top := monotonicStack[len(monotonicStack)-1]

			if nums[top] >= v {
				monotonicStack = append(monotonicStack, i%n)
				continue outer
			}
			ans[top] = v
			monotonicStack = monotonicStack[0 : len(monotonicStack)-1]
		}
	}
	return ans

}
