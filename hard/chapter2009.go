package hard

import (
	"math"
	"sort"

	"golang.org/x/exp/slices"
)

func MinOperations(nums []int) int {

	sort.Ints(nums)                      //排序
	distinctNums := slices.Compact(nums) //去重
	distinctNums = append(distinctNums, math.MaxInt32)
	min := math.MaxInt32

	for i := 0; i < len(distinctNums)-1; i++ {
		left := distinctNums[i]
		right := left + len(nums) - 1

		index := binarySearchLeft(distinctNums, right)

		operation := len(nums) - (index - i + 1)

		if min > operation {
			min = operation
		}
	}

	return min
}

// 大于target的左边界
func binarySearchLeft(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] <= target {
			l = mid + 1
		}

		if nums[mid] > target {
			r = mid
		}
	}
	return l - 1

}
