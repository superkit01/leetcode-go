package middle

func rearrangeArray(nums []int) []int {
	result := make([]int, len(nums))
	positive := 0
	negative := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result[positive] = nums[i]
			positive += 2
		} else {
			result[negative] = nums[i]
			negative += 2
		}
	}
	return result

}
