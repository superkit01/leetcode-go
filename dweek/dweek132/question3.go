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

func maximumLength(nums []int, k int) int {

	cache := make(map[[3]int]int, 0)

	var dfs func(index int, preValue int, diffCnt int) int

	dfs = func(index int, preValue int, diffCnt int) int {
		if v, ok := cache[[3]int{index, preValue, diffCnt}]; ok {
			return v
		}
		if diffCnt > k {
			cache[[3]int{index, preValue, diffCnt}] = math.MinInt32
			return math.MinInt32
		}

		if index >= len(nums) {
			cache[[3]int{index, preValue, diffCnt}] = 0
			return 0
		}
		value := 0
		//选择index位置
		if preValue == -1 { //初始值
			//选择index位置 （初始位置时 ums[index] 与 preValue不想等也不 +1 ）
			value = dfs(index+1, nums[index], diffCnt) + 1
			//不选择index位置
			value = max(dfs(index+1, preValue, diffCnt), value)

			cache[[3]int{index, preValue, diffCnt}] = value
			return value
		}

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
	return dfs(0, -1, 0)
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
