package week402

func countCompleteDayPairsII(hours []int) int64 {
	cache := make(map[int]int, 0)
	for _, v := range hours {
		cache[v%24]++
	}
	ans := int64(0)
	for k, v := range cache {
		if cache[(24-k)%24] == 0 {
			continue
		}
		if k > 12 {
			continue
		}
		if k == 0 || k == 12 {
			ans += int64(v) * int64((v - 1)) / 2
			continue
		}
		ans += int64(v) * int64(cache[(24-k)%24])
	}
	return ans
}
