package middle

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	dfs40(candidates, 0, target, make([]int, 0), &ans, true)

	return ans

}

func dfs40(candidates []int, index int, remaining int, current []int, ans *[][]int, use bool) {
	if remaining < 0 {
		return
	}

	if remaining == 0 {
		temp := make([]int, 0)
		temp = append(temp, current...)
		*ans = append(*ans, temp)
		return
	}

	if index >= len(candidates) {
		return
	}

	if !use && candidates[index] == candidates[index-1] {
		//如果当前位置与上一位相同，上一位没有使用，那么当前分支也不使用
		temp3 := make([]int, 0)
		temp3 = append(temp3, current...)
		dfs40(candidates, index+1, remaining, temp3, ans, false)
		return
	}

	//使用index位置
	temp2 := make([]int, 0)
	temp2 = append(temp2, current...)
	temp2 = append(temp2, candidates[index])
	dfs40(candidates, index+1, remaining-candidates[index], temp2, ans, true)

	//不使用index位置
	temp3 := make([]int, 0)
	temp3 = append(temp3, current...)

	dfs40(candidates, index+1, remaining, temp3, ans, false)

}
