package dweek132

import "math"

//暴力递归  TLE
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

//记忆化搜索
func MaximumLength(nums []int, k int) int {
	cache := make(map[[3]int]int, 0)

	var dfs func(index int, preValue int, diffCnt int) int

	dfs = func(index int, preValue int, diffCnt int) int {
		if diffCnt > k {
			return math.MinInt32
		}

		if index >= len(nums) {
			return 0
		}

		if v, ok := cache[[3]int{index, preValue, diffCnt}]; ok {
			return v
		}

		value := 0

		if nums[index] == preValue {
			//选择index位置
			value = dfs(index+1, nums[index], diffCnt) + 1
		} else {
			//选择index位置
			value = dfs(index+1, nums[index], diffCnt+1) + 1
			//不选择index位置
			value = max(dfs(index+1, preValue, diffCnt), value)
		}

		cache[[3]int{index, preValue, diffCnt}] = value
		return value

	}
	return dfs(0, -1, -1)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}


//DFS
func maximumLengthII(nums []int, k int) int {

	cache := make(map[[2]int]int, 0)

	//以index结尾 相邻不同数量为cnt 的最长好序列
	var dfs func(index int, diffCnt int) int
	dfs = func(index int, diffCnt int) int {

		if v, ok := cache[[2]int{index, diffCnt}]; ok {
			return v
		}

		if index == 0 {
			return 0
		}

		value := 0
		for j := index - 1; j >= 0; j-- {
			if nums[j] == nums[index] {
				value = max(value, dfs(j, diffCnt)+1)
			} else {
				if diffCnt-1 >= 0 {
					value = max(value, dfs(j, diffCnt-1)+1)
				}
			}
		}

		cache[[2]int{index, diffCnt}] = value
		return value

	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dfs(i, k))
	}
	return ans + 1
}
