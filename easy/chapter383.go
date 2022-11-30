package easy

func canConstruct(ransomNote string, magazine string) bool {

	cache := make(map[rune]int)

	for _, v := range magazine {
		if _, ok := cache[v]; !ok {
			cache[v] = 0
		}
		cache[v]++
	}

	for _, v := range ransomNote {
		if _, ok := cache[v]; !ok || cache[v] <= 0 {
			return false
		} else {
			cache[v]--
		}
	}
	return true

}
