package middle

func vowelStrings(words []string, queries [][]int) []int {
	vowel := []int{'a', 'e', 'i', 'o', 'u'}

	tmp := make([]int, 0)

	for _, e := range words {
		if contains(vowel, int(e[0])) && contains(vowel, int(e[len(e)-1])) {
			tmp = append(tmp, 1)
		} else {
			tmp = append(tmp, 0)
		}
	}

	if len(tmp) > 1 {
		for i := 1; i < len(tmp); i++ {
			tmp[i] = tmp[i-1] + tmp[i]
		}
	}

	result := make([]int, 0)

	for _, e := range queries {
		if e[0] == 0 {
			result = append(result, tmp[e[1]])
			continue

		}
		result = append(result, tmp[e[1]]-tmp[e[0]-1])

	}
	return result
}

// func contains(arr []int, t int) bool {
// 	for _, e := range arr {
// 		if e == t {
// 			return true
// 		}
// 	}
// 	return false
// }
