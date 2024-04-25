package lcr

func robII(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	//偷第一家
	dp := make([]int, len(nums)-1)
	dp[0] = nums[0]
	dp[1] = nums[0]
	for i := 2; i < len(nums)-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	//不偷第一家

	dp2 := make([]int, len(nums))
	dp2[0] = 0
	dp2[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}
	return max(dp[len(dp)-1], dp2[len(dp2)-1])

}
