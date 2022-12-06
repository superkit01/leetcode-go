package easy

func NumDifferentIntegers(word string) int {
	start := 0
	end := 0
	result := make(map[string]int, 0)
	for i, v := range word {

		if v >= 'a' && v <= 'z' {
			if start != end {
				value := word[start:end]
				for len(value) > 1 {
					if value[0] == '0' {
						value = value[1:]
					} else {
						break
					}
				}
				result[value] = 1
			}
			start = i + 1
			end = i + 1
		} else {
			end = i + 1
		}
	}

	if start != end {
		value := word[start:end]
		for len(value) > 1 {
			if value[0] == '0' {
				value = value[1:]
			} else {
				break
			}
		}
		result[value] = 1
	}

	return len(result)
}
