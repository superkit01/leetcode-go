package easy

import "math"

func rearrangeCharacters(s string, target string) int {
	targetMap := make(map[rune]int, 0)

	for _, v := range target {
		targetMap[v]++
	}

	resultCount := make(map[rune]int, 0)
	for _, v := range s {
		resultCount[v]++
	}
	min := math.MaxInt
	for k, v := range targetMap {
		if c, ok := resultCount[k]; ok {
			if c/v < min {
				min = c / v
			}
		} else {
			return 0
		}
	}
	return min
}
