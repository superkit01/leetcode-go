package lcr

import "sort"

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	ans := [][]int{}

	var dfs func(index int, current []int, target int, choose bool)

	dfs = func(index int, current []int, target int, choose bool) {

		if target == 0 || target < 0 {
			ans = append(ans, current)
			return
		}
		if index >= len(candidates) {
			return
		}

		if !choose && index > 0 && candidates[index] == candidates[index-1] {
			//剪枝
			//上一位置不选当前位置也不选
			dfs(index+1, current, target, false)
			return
		}
		//选择当前位置
		temp := []int{}
		temp = append(temp, current...)
		temp = append(temp, candidates[index])
		dfs(index, temp, target-candidates[index], true)

		temp2 := []int{}
		temp2 = append(temp2, current...)
		dfs(index+1, temp2, target, false)

	}

	dfs(0, []int{}, target, true)
	return ans

}
