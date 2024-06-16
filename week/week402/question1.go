package week402

import "fmt"

func countCompleteDayPairs(hours []int) int {
	cache := make(map[int]int, 0)
	for _, v := range hours {
		cache[v%24]++
	}
	ans := 0
	fmt.Printf("%v", cache)
	for k, v := range cache {
		if cache[(24-k)%24] == 0 {
			continue
		}
		if k > 12 {
			continue
		}
		if k == 0 || k == 12 {
			ans += v * (v - 1) / 2
			continue
		}
		ans += v * cache[(24-k)%24]
	}
	return ans
}
