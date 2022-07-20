package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countMap := make(map[rune]int)
	for _, v := range s {
		if _, ok := countMap[v]; !ok {
			countMap[v] = 0
		}
		countMap[v]++
	}

	for _, v := range t {
		if _, ok := countMap[v]; !ok {
			return false
		} else {
			countMap[v]--
			if countMap[v] < 0 {
				return false
			}
		}
	}
	return true

}
