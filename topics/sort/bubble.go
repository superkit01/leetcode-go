package sort
//冒泡
func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}

}
