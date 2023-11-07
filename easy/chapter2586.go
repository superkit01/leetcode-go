package easy

func vowelStrings(words []string, left int, right int) int {
	cache := make(map[byte]int, 0)
	cache['a'] = 0
	cache['e'] = 0
	cache['i'] = 0
	cache['o'] = 0
	cache['u'] = 0

	result := 0

	for i, v := range words {
		if i >= left && i <= right {
			if _, ok := cache[v[0]]; ok {
				if _, ok := cache[v[len(v)-1]]; ok {
					result += 1
				}

			}
		}
	}
	return result

}
