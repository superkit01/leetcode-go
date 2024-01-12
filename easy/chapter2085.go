package easy

func countWords(words1 []string, words2 []string) int {
	cache1 := make(map[string]int, 0)
	cache2 := make(map[string]int, 0)

	for _, v := range words1 {
		cache1[v]++
	}
	for _, v := range words2 {
		cache2[v]++
	}
	result := 0

	for k, v := range cache1 {
		if v == 1 && cache2[k] == 1 {
			result++
		}

	}

	return result

}
