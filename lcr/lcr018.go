package lcr

import "unicode"

func IsPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		for !(s[i] >= 'A' && s[i] <= 'Z') && !(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= '0' && s[i] <= '9') && i < j {
			i++
		}

		for !(s[i] >= 'A' && s[i] <= 'Z') && !(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= '0' && s[i] <= '9') && i < j {
			j--
		}
		if unicode.ToUpper(rune(s[i])) != unicode.ToUpper(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
