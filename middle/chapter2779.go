package middle

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	l := 0
	r := 0
	ans := 0
	for r < len(nums) {
		for r < len(nums) && nums[r]-nums[l] <= 2*k {
			r++
		}
		ans = max(r-l, ans)
		l++

	}
	return ans
}
