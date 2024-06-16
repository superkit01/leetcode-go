package week402

import (
	"sort"

	"golang.org/x/exp/slices"
)

//TLE
func maximumTotalDamage(power []int) int64 {
	damage := map[int]int64{}
	for _, v := range power {
		damage[v] += int64(v)
	}

	sort.Ints(power)
	power = slices.Compact(power)

	memo := map[[2]int]int64{}

	var dfs func(index int, edge int) int64
	dfs = func(index int, edge int) int64 {

		if v, ok := memo[[2]int{index, edge}]; ok {
			return v
		}

		if index >= len(power) {
			return 0
		}
		value := int64(0)
		if power[index] <= edge {
			//不选index
			value = dfs(index+1, edge)
		} else {
			//选择index
			value = dfs(index+1, power[index]+2) + damage[power[index]]
			//不选index
			value = max(value, dfs(index+1, edge))
		}
		memo[[2]int{index, edge}] = value
		return value
	}
	return dfs(0, -1)

}

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

func MaximumTotalDamageII(power []int) int64 {
	damage := map[int]int64{}
	for _, v := range power {
		damage[v] += int64(v)
	}
	sort.Ints(power)
	power = slices.Compact(power)

	memo := map[int]int64{}

	var dfs func(index int) int64
	dfs = func(index int) int64 {

		if v, ok := memo[index]; ok {
			return v
		}

		if index < 0 {
			return int64(0)
		}

		value := dfs(index - 1)

		j := index
		for i := index; i >= 0; i-- {
			if power[i]+2 < power[index] {
				j = i
				break
			}
		}
		if j == index { //没找到
			value = max(value, damage[power[index]])
		} else {
			value = max(value, dfs(j)+damage[power[index]])
		}

		memo[index] = value
		return value
	}
	return dfs(len(power) - 1)

}
