package middle

func validPartition(nums []int) bool {
	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return true
		} else {
			return false
		}
	}
	// 1 1 1 3 3  5 5 5  7 7 7
	// F T T 
	ans := make([]bool, len(nums))

	ans[0] = false
	ans[1] = nums[0] == nums[1]
	ans[2] = (nums[0] == nums[1] && nums[1] == nums[2]) || (nums[2]-nums[1] == 1 && nums[1]-nums[0] == 1)
	if len(nums) == 3 {
		return ans[2]
	}

	for i := 3; i < len(ans); i++ {
		if ans[i-2] {
			if nums[i] == nums[i-1] {
				ans[i] = true
				continue
			}

		}

		if i-3 < 0 {
			if (nums[i-1]-nums[i-2] == 1 && nums[i]-nums[i-1] == 1) || (nums[i] == nums[i-1] && nums[i-1] == nums[i-2]) {
				ans[i] = true
				continue
			}

		} else {
			if ans[i-3] {
				if (nums[i-1]-nums[i-2] == 1 && nums[i]-nums[i-1] == 1) || (nums[i] == nums[i-1] && nums[i-1] == nums[i-2]) {
					ans[i] = true
					continue
				}
			}
		}
		ans[i] = false
	}

	return ans[len(ans)-1]

}
