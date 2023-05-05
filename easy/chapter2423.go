package easy

func equalFrequency(word string) bool {
	cache := make(map[rune]int, 0)
	for _, v := range word {
		cache[v]++
	}

	if len(cache) == 1 {
		return true
	}

	count := make([]int, 0)
	for _, v := range cache {
		count = append(count, v)
	}

	for i := 0; i < len(count); i++ {
		count[i]--

		temp := make([]int, 0)
		for _, v := range count {
			if v != 0 {
				temp = append(temp, v)
			}
		}

		if checkCount(temp) {
			return true
		}

		count[i]++

	}
	return false

}

func checkCount(count []int) bool {
	if len(count) == 1 {
		return true
	}
	for i := 0; i < len(count)-1; i++ {
		if count[i] != count[i+1] {
			return false
		}
	}
	return true
}
