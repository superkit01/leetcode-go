package week394

func numberOfSpecialCharsII(word string) int {
	offset := int('A') - int('a')

	cache := make(map[int]int, 0)

	for i, v := range word {
		if v >= 'a' && v <= 'z' {
			cache[int(v)] = i
		}
	}

	ans := 0
	cache2 := make(map[int]bool, 0)

	for i, v := range word {
		if v >= 'A' && v <= 'Z' {
			lowerCase := int(v) - offset
			if v, ok := cache[lowerCase]; ok {
				if v >= i {
					cache2[lowerCase] = true
				} else {
					if _, ok := cache2[lowerCase]; !ok {
						ans += 1
						cache2[lowerCase] = true
					}
				}
			}
		}
	}
	return ans
}
