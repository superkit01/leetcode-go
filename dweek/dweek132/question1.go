package dweek132

func clearDigits(s string) string {
	ans := []rune{}

	for _, v := range s {

		if v >= '0' && v <= '9' {
			if len(ans) > 0 {
				ans = ans[0 : len(ans)-1]
			}
			continue
		}
		ans = append(ans, v)
	}

	res := ""
	for _, v := range ans {
		res += string(v)
	}
	return res
}
