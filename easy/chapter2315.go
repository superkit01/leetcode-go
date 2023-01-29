package easy

func countAsterisks(s string) int {
	flag := true
	count := 0
	for _, v := range s {
		if v == '|' {
			flag = !flag
		}
		if flag && v == '*' {
			count++
		}

	}
	return count
}
