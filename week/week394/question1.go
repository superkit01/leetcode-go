package week394

func numberOfSpecialChars(word string) int {
	offset := int('A') - int('a')

	cache := make(map[int]bool, 0)
	for _, v := range word {
		if v >= 'A' && v <= 'Z' {
			cache[int(v)] = true
		}
	}

	ans := 0
	cache2 := make(map[int]bool, 0)
	for _, v := range word {
		if v >= 'a' && v <= 'z' {
			lowerCase := int(v) + offset
			if _, ok := cache[lowerCase]; ok {
				if _, ok := cache2[lowerCase]; !ok {
					ans += 1
					cache2[lowerCase] = true
				}
			}
		}
	}

	return ans
}
