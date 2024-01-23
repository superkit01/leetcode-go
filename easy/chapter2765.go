package easy

func alternatingSubarray(nums []int) int {

	cache := make([]int, len(nums))
	cache[0] = 1
	if nums[1]-nums[0] == 1 {
		cache[1] = 2
	} else {
		cache[1] = 1
	}
	if len(nums) == 2 {
		if cache[1] == 2 {
			return 2
		} else {
			return -1
		}

	}

	for i := 2; i < len(nums); i++ {
		if cache[i-1] != 1 && nums[i] == nums[i-2] {
			cache[i] = cache[i-1] + 1
		} else if nums[i]-nums[i-1] == 1 {
			cache[i] = 2
		} else {
			cache[i] = 1
		}
	}

	max := 0
	for i := 0; i < len(cache); i++ {
		if max < cache[i] {
			max = cache[i]
		}
	}

	if max == 1 {
		return -1
	} else {
		return max
	}

}
