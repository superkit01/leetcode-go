package lcr

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			left += nums[i-1]
		}
		sum -= nums[i]
		if left == sum {
			return i
		}
	}
	return -1
}
