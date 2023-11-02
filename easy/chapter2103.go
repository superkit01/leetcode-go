package easy

func CountPoints(rings string) int {
	container := make(map[byte]int, 0)

	for i := 1; i < len(rings); i += 2 {
		if _, ok := container[rings[i]]; !ok {
			container[rings[i]] = 0
		}
		if rings[i-1] == 'B' {
			container[rings[i]] = container[rings[i]] | 0b100
		}
		if rings[i-1] == 'G' {
			container[rings[i]] = container[rings[i]] | 0b010
		}
		if rings[i-1] == 'R' {
			container[rings[i]] = container[rings[i]] | 0b001
		}

	}

	result := 0

	for _, v := range container {
		if v&0b111 == 0b111 {
			result++
		}
	}

	return result

}
