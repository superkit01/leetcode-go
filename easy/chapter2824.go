package easy

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	result := 0

	for i := 0; i < len(nums)-1; i++ {
		minus := target - nums[i]

		if nums[i+1] >= minus {
			continue
		}
		if nums[len(nums)-1] < minus {
			result += len(nums) - 1 - i
			continue
		}
		index := binarySearch(i, nums, target-nums[i])
		result += index - i

	}
	return result
}

func binarySearch(i int, nums []int, target int) int {
	left := i
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1

}
