package easy

import (
	"strconv"
	"strings"
	"unicode"
)

func AreNumbersAscending(s string) bool {
	tokens := strings.Split(s, " ")

	result := make([]int, 0)
outer:
	for _, token := range tokens {
		for _, d := range token {
			if !unicode.IsDigit(d) {
				continue outer
			}
		}
		value, _ := strconv.Atoi(token)
		result = append(result, value)
	}

	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true

}
