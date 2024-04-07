package week392

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	maxUp := 1
outer:
	for i := 1; i < len(nums); i++ {
		up := 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[j+1] {
				up++
				if maxUp < up {
					maxUp = up
				}
			} else {
				continue outer
			}
		}
	}

	maxDown := 1
second:
	for i := 1; i < len(nums); i++ {
		down := 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				down++
				if maxDown < down {
					maxDown = down
				}
			} else {
				continue second
			}

		}

	}

	if maxUp > maxDown {
		return maxUp
	} else {
		return maxDown
	}

}
