package middle

func numMatchingSubseq(s string, words []string) int {
	result := 0
	cache := make(map[string]bool)
	for _, v := range words {
		if b, ok := cache[v]; ok && b {
			result++
			continue
		}

		if match(s, v) {
			cache[v] = true
			result++
		} else {
			cache[v] = false
		}
	}
	return result
}

func match(s string, v string) bool {
	i := 0
	j := 0

	for {
		if j == len(v) {
			return true
		}
		if i == len(s) {
			return false
		}

		if v[j] == s[i] {
			i++
			j++
		} else {
			i++
		}

	}
}
