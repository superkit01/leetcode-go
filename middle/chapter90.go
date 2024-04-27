package middle

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	dfs90(nums, 0, make([]int, 0), &ans, true)
	return ans
}

func dfs90(nums []int, index int, current []int, ans *[][]int, use bool) {
	if index == len(nums) {
		*ans = append(*ans, current)
		return
	}

	if index > 0 && nums[index] == nums[index-1] && !use {
		temp := make([]int, 0)
		temp = append(temp, current...)
		dfs90(nums, index+1, temp, ans, false)
		return
	}

	//使用当前位置
	temp2 := make([]int, 0)
	temp2 = append(temp2, current...)
	temp2 = append(temp2, nums[index])
	dfs90(nums, index+1, temp2, ans, true)

	//不使用当前位置
	temp3 := make([]int, 0)
	temp3 = append(temp3, current...)
	dfs90(nums, index+1, temp3, ans, false)
}
