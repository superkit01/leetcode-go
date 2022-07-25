package main

func firstUniqChar(s string) int {
	cache := make(map[rune]int)

	for _, v := range s {
		if _, ok := cache[v]; !ok {
			cache[v] = 0
		}
		cache[v]++
	}

	for i, v := range s {
		if cache[v] == 1 {
			return i
		}
	}
	return -1
}
