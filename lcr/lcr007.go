package lcr

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	if nums[0] > 0 {
		return [][]int{}
	}

	ans := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
	middle:
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue middle
			}

			target := -nums[i] - nums[j]
			index := binarySearch(nums, j+1, len(nums)-1, target)
			if index != -1 {
				ans = append(ans, []int{nums[i], nums[j], nums[index]})
			}

		}
	}

	return ans
}

func binarySearch(nums []int, l int, r int, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1

}
