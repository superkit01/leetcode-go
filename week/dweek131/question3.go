package dweek131

func QueryResults(limit int, queries [][]int) []int {
	ori := map[int]int{}
	cache := make(map[int]int, 0)
	ans := make([]int, len(queries))
	for i, v := range queries {

		pre := ori[v[0]]
		if _, ok := cache[pre]; ok {
			if cache[pre] == 1 {
				delete(cache, pre)
			} else {
				cache[pre]--
			}
		}

		cache[v[1]]++
		ori[v[0]] = v[1]

		ans[i] = len(cache)
	}
	return ans

}
