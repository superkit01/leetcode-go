package hard

func minDays(n int) int {
	cache := make(map[int]int, 0)
	return dfs(n, &cache)

}

func dfs(n int, cache *map[int]int) int {
	if v, ok := (*cache)[n]; ok {
		return v
	}
	if n <= 2 {
		(*cache)[n] = n
		return n
	}
	temp := min(dfs(n/2, cache)+1+n%2, dfs(n/3, cache)+1+n%3)
	(*cache)[n] = temp
	return temp

}

