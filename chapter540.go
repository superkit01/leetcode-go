package main



func singleNonDuplicate(nums []int) int {
	if len(nums)%2 == 0 {
		return 0
	}
	tmp := 0
	//for _, v := range nums {
	//	tmp = v ^ tmp
	//}

	for i := 0; i <= len(nums)/2; i++ {
		if i <len(nums)/2 {
			tmp = nums[i] ^ tmp
			tmp = nums[len(nums)-1-i]
		} else {
			tmp = nums[i] ^ tmp
		}
		return tmp

	}

	return tmp
}
