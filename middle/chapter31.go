package middle

import "sort"

func nextPermutation(nums []int) {

	tempIndex := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			tempIndex = i
			break
		}
	}

	if tempIndex == 0 {
		sort.Ints(nums)
		return
	}
	for i := len(nums) - 1; i >= tempIndex; i-- {
		if nums[i] > nums[tempIndex-1] {
			temp := nums[tempIndex-1]
			nums[tempIndex-1] = nums[i]
			nums[i] = temp
			break
		}
	}
	sort.Ints(nums[tempIndex:])

}
