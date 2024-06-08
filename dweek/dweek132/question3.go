package dweek132

func MaximumLengthII(nums []int, k int) int {
	maxV := 0
	var dfs func(index int, preValue int, cnt int, length int)

	dfs = func(index int, preValue int, cnt int, length int) {
		if cnt > k {
			return
		}
		maxV = max(maxV, length)

		if index >= len(nums) {
			return
		}

		//选择当前位置
		if preValue == -1 {
			if nums[index] != preValue {
				dfs(index+1, nums[index], cnt, length+1)
			} else {
				dfs(index+1, nums[index], cnt, length+1)
			}
		} else {
			if nums[index] != preValue {
				dfs(index+1, nums[index], cnt+1, length+1)
			} else {
				dfs(index+1, nums[index], cnt, length+1)
			}

		}

		//不选择当前位置
		dfs(index+1, preValue, cnt, length)
	}

	dfs(0, -1, 0, 0)
	return maxV

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
