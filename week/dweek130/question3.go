package dweek130

import "math"

func minimumSubstringsInPartition(s string) int {
	//记忆化搜索
	cache := make(map[string]int, 0)
	return dfs(s, &cache) + 1
}

func dfs(s string, cache *map[string]int) int {
	if v, ok := (*cache)[s]; ok {
		return v
	}
	//Break Case
	if len(s) <= 1 {
		(*cache)[s] = 0
		return 0
	}
	if isBalance(s) {
		(*cache)[s] = 0
		return 0
	}

	value := math.MaxInt32
	//不平衡遍历切分
	for i := 1; i < len(s); i++ {
		temp := dfs(s[0:i], cache) + dfs(s[i:], cache) + 1
		if value > temp {
			value = temp
		}
	}
	(*cache)[s] = value
	return value
}

func isBalance(s string) bool {
	if len(s) <= 1 {
		return true
	}

	cnt := [26]int{}
	for _, v := range s {
		cnt[v-'a']++
	}

	flag := false
	value := 0
	for _, v := range cnt {
		if v != 0 {
			if !flag {
				value = v
				flag = !flag
			}
			if v != value {
				return false
			}
		}

	}
	return true
}
