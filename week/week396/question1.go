package week396

import "golang.org/x/exp/slices"

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	cache := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	aFlag := false
	bFlag := false
	for _, v := range word {
		if v == '@' || v == '#' || v == '$' {
			return false
		}

		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			if slices.Contains(cache, v) {
				aFlag = true
			} else {
				bFlag = true
			}
		}
	}
	return aFlag && bFlag

}
