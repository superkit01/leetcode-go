package middle

func canJump(nums []int) bool {

outer:
	for i := len(nums) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if j+nums[j] >= i {
				continue outer
			}
		}
		return false
	}
	return true

}
