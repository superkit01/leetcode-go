package template

// 1, 2, 2, 3, 4, 7, 7, 7, 7, 7, 8, 9, 10, 10, 10, 18

// 7 小于等于7 的最大值
func UpperBound(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1

		if nums[mid] <= target {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 7 大于等于7 的最小值
func LowerBound(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1

		if nums[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
