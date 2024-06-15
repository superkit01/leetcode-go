package lcr

//DFS
func combinationSum4(nums []int, target int) int {

	memo := map[int]int{}

	var dfs func(target int) int

	dfs = func(target int) int {
		if _, ok := memo[target]; ok {
			return memo[target]
		}
		if target == 0 {
			memo[target] = 1
			return 1
		}

		if target < 0 {
			memo[target] = 0
			return 0
		}

		value := 0

		for i := 0; i < len(nums); i++ {
			value += dfs(target - nums[i])
		}
		memo[target] = value
		return value

	}
	return dfs(target)
}

//DP
func combinationSum4II(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1

	for i := range dp {
		for _, v := range nums {
			if i-v >= 0 {
				dp[i] = dp[i] + dp[i-v]
			}
		}
	}
	return dp[target]

}
