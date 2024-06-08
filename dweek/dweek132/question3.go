package dweek132

//TLE
func MaximumLengthII(nums []int, k int) int {
	maxV := 0

	var dfs func(index int, preValue int, diffCnt int, length int)

	dfs = func(index int, preValue int, diffCnt int, length int) {
		if diffCnt > k {
			return
		}
		maxV = max(maxV, length)

		if index >= len(nums) {
			return
		}

		//选择index位置
		if preValue == -1 { //初始值
			dfs(index+1, nums[index], diffCnt, length+1)
		} else {
			if nums[index] != preValue {
				dfs(index+1, nums[index], diffCnt+1, length+1)
			} else {
				dfs(index+1, nums[index], diffCnt, length+1)
			}
		}

		//不选择index位置
		dfs(index+1, preValue, diffCnt, length)

	}

	dfs(0, -1, 0, 0)

	return maxV

}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
