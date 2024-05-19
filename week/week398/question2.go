package week398

func isArraySpecialII(nums []int, queries [][]int) []bool {
	dp := make([]int, len(nums)) //以i结尾的最长特殊数组长度
	dp[0] = 1
	if len(dp) > 1 {
		for i := 1; i < len(dp); i++ {
			if nums[i]%2 != nums[i-1]%2 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = 1
			}
		}
	}
	ans := make([]bool, len(queries))
	for i, v := range queries {
		if v[1]-v[0]+1 <= dp[v[1]] {
			ans[i] = true
		}
	}
	return ans

}
