package easy

func distributeCandies(n int, limit int) int {
	ans := 0
	for i := 0; i <= n && i <= limit; i++ {
		for j := 0; j <= n-i && j <= limit; j++ {
			if n-i-j <= limit {
				ans++
			}
		}
	}
	return ans
}
