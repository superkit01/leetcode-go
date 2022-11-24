package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	for i, v := range nums {
		if v < left {
			nums[i] = 0
		} else if v <= right {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}

	last1 := -1
	last2 := -1
	result := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			last1 = i
		}
		if nums[i] == 2 {
			last2 = i
			last1 = -1
		}

		if last1 != -1 {
			result += last1 - last2
		}

	}
	return result

}
