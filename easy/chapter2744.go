package easy

func maximumNumberOfStringPairs(words []string) int {
	cache := make(map[string]int, 0)

	for i, v := range words {
		cache[v] = i
	}

	result := 0

	for i, v := range words {
		if index, ok := cache[reverseWord(v)]; ok {
			if index != i {
				result++
			}

		}
	}
	return result / 2

}

func reverseWord(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
