package middle

func maxArrayValue(nums []int) int64 {
	var maxValue int64 = int64(nums[len(nums)-1])

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] >= nums[i-1] {
			nums[i-1] = nums[i] + nums[i-1]
		}

		if maxValue < int64(nums[i-1]) {
			maxValue = int64(nums[i-1])
		}

	}

	return maxValue

}
