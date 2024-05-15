package lcr

func findTargetSumWays(nums []int, target int) int {

	memo := make(map[int]map[int]int, 0)

	var dfs func(index, rest int) int
	dfs = func(index, rest int) int {
		if v, ok := memo[index]; ok {
			if m, ok := v[rest]; ok {
				return m
			}
		} else {
			memo[index] = make(map[int]int, 0)
		}

		if index == -1 {

			if rest == 0 {
				memo[index][rest] = 1
				return 1
			} else {
				memo[index][rest] = 0
				return 0
			}
		}
		value := dfs(index-1, rest-nums[index]) + dfs(index-1, rest+nums[index])
		memo[index][rest] = value
		return value
	}

	return dfs(len(nums)-1, target)

}

// 添加+号 的和为 positiveSum
// 添加-号 的和为 negativeSum
// positiveSum-negativeSum=target
// (totalSum-negativeSum)-negativeSum=target
// negativeSum=(totalSum-target)/2
func findTargetSumWaysII(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum < target || (sum-target)&0b1 == 0b1 {
		return 0
	}

	neg := (sum - target) >> 1

	// 原题转化为从nums中选取n个元素和为neg的方案数

	//前i个元素和为j的方案数量
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1

	for index := range nums {
		for j := 0; j <= neg; j++ {
			if j >= nums[index] {
				dp[index+1][j] = dp[index][j] + dp[index][j-nums[index]]
			} else {
				dp[index+1][j] = dp[index][j]
			}
		}
	}

	return dp[len(dp)-1][neg]

}
