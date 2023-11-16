package easy

import "math"

func longestAlternatingSubarray(nums []int, threshold int) int {
	mod := 0
	start := -1
	max := 0
	for i, v := range nums {
		if start == -1 {
			if v%2 == 0 && v <= threshold {
				start = i
				mod = 1
			}
			continue
		}

		if v%2 == mod && v <= threshold {
			mod = ^mod & 0b01
		} else {
			max = int(math.Max(float64(i-start), float64(max)))
			if v%2 == 0 && v <= threshold {
				start = i
				mod = 1
			} else {
				start = -1
				mod = 0
			}
		}

	}

	if start != -1 {
		max = int(math.Max(float64(len(nums)-start), float64(max)))
	}
	return max

}
