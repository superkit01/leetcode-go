package easy

func makeSmallestPalindrome(s string) string {
	r := make([]byte, len(s), len(s))
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			if s[i] < s[j] {
				r[i] = s[i]
				r[j] = s[i]
			} else {
				r[i] = s[j]
				r[j] = s[j]
			}
		} else {
			r[i] = s[i]
			r[j] = s[i]
		}

		i++
		j--

	}
	return string(r)

}
