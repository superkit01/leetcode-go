package middle

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1Map := make(map[rune]int, 0)
	for _, v := range word1 {
		word1Map[v]++
	}

	word2Map := make(map[rune]int, 0)
	for _, v := range word2 {
		word2Map[v]++
	}

	if len(word1Map) != len(word2Map) {
		return false
	}
	for k, _ := range word1Map {
		if _, ok := word2Map[k]; !ok {
			return false
		}
	}

	word1CountMap := make(map[int]int, 0)
	for _, v := range word1Map {
		word1CountMap[v]++
	}

	word2CountMap := make(map[int]int, 0)
	for _, v := range word2Map {
		word2CountMap[v]++
	}

	for k, v := range word1CountMap {
		if v != word2CountMap[k] {
			return false
		}

	}

	return true
}
