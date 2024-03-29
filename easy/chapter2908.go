package easy

import "math"

func MinimumSum(nums []int) int {
	prefixMin := make([]int, len(nums))
	prefixMin[0] = nums[0]
	suffixMin := make([]int, len(nums))
	suffixMin[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		if nums[i] < prefixMin[i-1] {
			prefixMin[i] = nums[i]
		} else {
			prefixMin[i] = prefixMin[i-1]
		}

		if nums[len(nums)-1-i] < suffixMin[len(nums)-i] {
			suffixMin[len(nums)-1-i] = nums[len(nums)-1-i]
		} else {
			suffixMin[len(nums)-1-i] = suffixMin[len(nums)-i]
		}
	}

	min := math.MaxInt32
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] <= prefixMin[i-1] {
			continue
		}

		if nums[i] <= suffixMin[i+1] {
			continue
		}

		if min > nums[i]+prefixMin[i-1]+suffixMin[i+1] {
			min = nums[i] + prefixMin[i-1] + suffixMin[i+1]

		}
	}
	if min == math.MaxInt32 {
		return -1
	} else {
		return min
	}

}
