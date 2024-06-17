package week402

func countCompleteDayPairsII(hours []int) int64 {
	cache := make(map[int]int, 0)
	ans := int64(0)
	for _, v := range hours {
		ans += int64(cache[(24-v%24)%24])
		cache[v%24]++
	}
	return ans
}
