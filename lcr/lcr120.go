package lcr

func findRepeatDocument(documents []int) int {
	hash := make(map[int]int, 0)

	for _, v := range documents {
		if _, ok := hash[v]; ok {
			return v
		}
		hash[v] = v
	}
	return -1

}
