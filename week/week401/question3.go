package week401

import (
	"sort"

	"golang.org/x/exp/slices"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	rewardValues = slices.Compact(rewardValues)

	cache := make(map[[2]int]int, 0)

	var dfs func(index int, current int) int

	dfs = func(index int, current int) int {
		if v, ok := cache[[2]int{index, current}]; ok {
			return v
		}

		if index >= len(rewardValues) {
			cache[[2]int{index, current}] = current
			return current
		}
		ans := 0
		if rewardValues[index] <= current {
			//不取当前位置
			ans = max(dfs(index+1, current), ans)
		} else {
			//取当前位置
			ans = max(dfs(index+1, current+rewardValues[index]), ans)
			//不取当前位置
			ans = max(dfs(index+1, current), ans)
		}
		cache[[2]int{index, current}] = current
		return ans
	}

	return dfs(0, 0)

}
