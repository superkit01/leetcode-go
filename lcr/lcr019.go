package lcr

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		return valid(s, i+1, j) || valid(s, i, j-1)
	}
	return true

}

func valid(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--

	}
	return true

}
