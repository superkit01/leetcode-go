package easy

func longestPalindrome(s string) int {
	cache := make(map[rune]int, 0)
	for _, v := range s {
		cache[v]++
	}

	result := 0
	flag := false
	for _, v := range cache {
		if v%2 == 0 {
			result += v
		} else {
			result += (v - 1)
			flag = true
		}
	}
	if flag {
		result += 1
	}
	return result

}
