package middle

func minSubArrayLen(target int, nums []int) int {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = nums[i] + prefix[i]
	}
	if prefix[len(prefix)-1] < target {
		return 0
	}
	ans := len(nums)

	i := 0
	j := 0
	for j < len(prefix) && i < len(prefix) {
		if prefix[j]-prefix[i] >= target {
			if ans > j-i {
				ans = j - i
			}
			i++
			continue
		}
		j++
	}
	return ans

}
