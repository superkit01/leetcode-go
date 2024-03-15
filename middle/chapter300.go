package middle

//1,3,6,7,9,4
func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		max := 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if max < dp[j]+1 {
					max = dp[j] + 1
				}

			}
		}
		dp[i] = max
	}

	ans := 0
	for _, v := range dp {
		if ans < v {
			ans = v
		}
	}
	return ans

}
