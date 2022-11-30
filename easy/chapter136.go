package easy

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp ^= nums[i]
	}
	return tmp

}
