package main

func moveZeroes(nums []int) {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			continue
		} else {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}

}
