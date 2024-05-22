package middle

func productExceptSelf(nums []int) []int {
	suffix := make([]int, len(nums))
	suffix[len(suffix)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}
	prefix := 1
	ans := make([]int, len(nums))

	for i := range ans {
		if i == len(ans)-1 {
			ans[i] = prefix
			continue
		}
		ans[i] = prefix * suffix[i+1]
		prefix *= nums[i]
	}
	return ans

}
