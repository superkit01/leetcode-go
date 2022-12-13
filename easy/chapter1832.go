package easy

func checkIfPangram(sentence string) bool {
	cache := make(map[int]bool, 0)

	for _, v := range sentence {
		cache[int(v-'a')] = true
	}

	for i := 0; i < 26; i++ {
		if _, ok := cache[i]; !ok {
			return false
		}
	}
	return true

}
