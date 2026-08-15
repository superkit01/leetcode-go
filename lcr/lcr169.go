package lcr

func dismantlingActionII(arr string) byte {
	s := make([]int, 26)
	for _, v := range arr {
		s[v-'a']++
	}

	for _, v := range arr {
		if s[v-'a'] == 1 {
			return byte(v)
		}
	}

	return ' '
}
