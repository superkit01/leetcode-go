package dweek130

import "math"

func MinimumSubstringsInPartitionII(s string) int {
	//记忆化搜索
	cache := make(map[int]int, 0)

	return dfsII(s, len(s)-1, &cache) + 1
}

func dfsII(s string, index int, cache *map[int]int) int {
	if v, ok := (*cache)[index]; ok {
		return v
	}

	if index < 0 {
		(*cache)[index] = 0
		return 0
	}

	cnt := [26]int{}
	maxCount := 0
	maxLetters := 0
	value := math.MaxInt

	for i := index; i >= 0; i-- {
		if cnt[s[i]-'a'] == 0 {
			maxLetters++
		}
		cnt[s[i]-'a']++
		if maxCount < cnt[s[i]-'a'] {
			maxCount = cnt[s[i]-'a']
		}
		if maxLetters*maxCount == index-i+1 {
			temp := dfsII(s, i-1, cache) + 1
			if value > temp {
				value = temp
			}
		}
	}
	(*cache)[index] = value
	return value
}
