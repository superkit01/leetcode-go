package lcr

func NumSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	i := 0
	j := 0
	multiply := 1
	for j < len(nums) {
		multiply *= nums[j]
		for multiply >= k && i <= j {
			multiply /= nums[i]
			i++
		}
		ans += j - i + 1
		j++
	}
	return ans

}
