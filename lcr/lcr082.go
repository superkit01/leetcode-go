package lcr

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := [][]int{}

	var dfs func(index int, current []int, target int, choose bool)
	dfs = func(index int, current []int, target int, choose bool) {
		if target == 0 {
			ans = append(ans, current)
			return
		}
		if index >= len(candidates) || target < 0 {
			return
		}

		if index > 0 && candidates[index] == candidates[index-1] && !choose {
			//不选择当前位置
			dfs(index+1, current, target, choose)
			return
		}

		//选择当前位置
		temp := append(append([]int{}, candidates[index]), current...)
		dfs(index+1, temp, target-candidates[index], true)

		//不选择当前位置
		temp2 := append([]int{}, current...)
		dfs(index+1, temp2, target, false)

	}

	dfs(0, []int{}, target, true)

	return ans

}
