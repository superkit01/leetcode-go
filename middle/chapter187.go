package middle

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	ans := []string{}

	cache := make(map[string]int, 0)

	for i := 0; i < len(s)-9; i++ {
		temp := s[i : i+10]

		cache[temp]++
		if cache[temp] == 2 {
			ans = append(ans, temp)
		}
	}
	return ans
}
