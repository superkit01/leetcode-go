package lcr

import "math"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		} else {
			return 0
		}
	}

	prefixSum := make([]int, len(nums))

	for i := 0; i < len(prefixSum); i++ {
		if nums[i] >= target {
			return 1
		}
		if i == 0 {
			prefixSum[i] = nums[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + nums[i]
		}
	}

	if prefixSum[len(prefixSum)-1] < target {
		return 0
	}

	ans := math.MaxInt32
	for i := 0; i < len(prefixSum); i++ {
		if prefixSum[i] >= target {
			index := binarySearchII(prefixSum, prefixSum[i]-target, i)
			if ans > i-index {
				ans = i - index
			}
		}
	}
	return ans

}

// 小于等于target 的最大位置
func binarySearchII(prefixSum []int, target int, i int) int {
	l := 0
	r := i

	for l < r {
		mid := (l + r) / 2

		if prefixSum[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l - 1

}
