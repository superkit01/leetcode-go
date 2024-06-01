package middle

func distributeCandies(n int, limit int) int64 {
	ans := int64(0)
	for i := 0; i <= n && i <= limit; i++ {
		if n-i > limit*2 {
			continue
		}
		//100   55
		// ans+= int64(min(n-i,limit) - max(0,n-i-limit)+1)

		ans += int64(min(n-i, limit) - (n - i - (min(n-i, limit))) + 1)
	}
	return ans
}
