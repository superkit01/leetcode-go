package easy

func applyOperations(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
		}
	}

	slow := 0
	fast := 0
outer:
	for slow < len(nums)-1 {
		if nums[slow] != 0 {
			slow++
			fast++
			continue
		}
	inner:
		for {
			if slow == fast {
				fast++
				continue inner
			}
			if fast >= len(nums) {
				break outer
			}
			if nums[fast] != 0 {
				nums[slow], nums[fast] = nums[fast], nums[slow]
				slow++
				continue outer
			} else {
				fast++
			}

		}
	}
	return nums

}
