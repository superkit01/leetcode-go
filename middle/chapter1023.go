package middle

import (
	"unicode"
)

func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))
outer:
	for k, querie := range queries {
		result[k] = true
		i := 0
		for _, v := range querie {
			if i < len(pattern) && v == rune(pattern[i]) {
				i++
			} else if unicode.IsUpper(v) {
				result[k] = false
				continue outer
			}
		}
		if i < len(pattern) {
			result[k] = false
		}

	}
	return result

}
