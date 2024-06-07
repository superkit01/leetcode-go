package easy

func maxOperations(nums []int) int {
	if len(nums) == 2 {
		return 1
	}
	var dfs func(i, j, score int) int
	dfs = func(i, j, score int) int {
		if i >= j {
			return 0
		}
		ans := 0
		if nums[i]+nums[i+1] == score {
			ans = max(dfs(i+2, j, score)+1, ans)
		}
		return ans
	}

	return dfs(0, len(nums)-1, nums[0]+nums[1])

}
