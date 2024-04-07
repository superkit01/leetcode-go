package week392

import "sort"

func MinOperationsToMakeMedianK(nums []int, k int) int64 {
	var ans int64

	sort.Ints(nums)

	pos := len(nums) / 2

	for i := 0; i < len(nums); i++ {
		if i < pos {
			if nums[i] > k {
				ans += int64(nums[i] - k)
			}
		}

		if i == pos {
			if nums[i] > k {
				ans += int64(nums[i] - k)
			} else {
				ans += int64(k - nums[i])
			}
		}

		if i > pos {
			if nums[i] < k {
				ans += int64(k - nums[i])
			}
		}
	}
	return ans

}
