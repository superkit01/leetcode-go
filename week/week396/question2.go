package week396

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	cache := make(map[string]int)
	for i := 0; i < len(word)/k; i++ {
		v := word[k*i : k*i+k]
		cache[v]++
	}
	max := 0
	ans := 0
	for _, v := range cache {
		if max < v {
			max = v
		}
		ans += v
	}
	return ans - max

}
