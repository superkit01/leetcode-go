package main

import "strings"

func wordPattern(pattern string, s string) bool {

	sSlice := strings.Split(s, " ")
	patternSlice := strings.Split(pattern, "")
	if len(sSlice) != len(patternSlice) {
		return false
	}

	map1 := make(map[string]string)
	map2 := make(map[string]string)

	for i := 0; i < len(patternSlice); i++ {
		if _, ok := map1[patternSlice[i]]; !ok {
			map1[patternSlice[i]] = sSlice[i]
		}

		if map1[patternSlice[i]] != sSlice[i] {
			return false
		}

		if _, ok := map2[sSlice[i]]; !ok {
			map2[sSlice[i]] = patternSlice[i]
		}

		if map2[sSlice[i]] != patternSlice[i] {
			return false
		}
	}
	return true
}

func main() {
	wordPattern("abba", "cat dog dog cat")
}
