package lcr

func subsets(nums []int) [][]int {
	ans := [][]int{}

	var dfs func(i int, current []int)
	dfs = func(i int, current []int) {
		if i == len(nums) {
			ans = append(ans, current)
			return
		}
		//不选当前位置
		temp2 := []int{}
		temp2 = append(temp2, current...)
		dfs(i+1, temp2)

		//选择当前位置
		temp := []int{}
		temp = append(temp, current...)
		temp = append(temp, nums[i])
		dfs(i+1, temp)

	}
	dfs(0, []int{})
	return ans

}
