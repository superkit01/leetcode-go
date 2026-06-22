package week400

func minimumChairs(s string) int {

	ans := 0
	current := 0

	for _, v := range s {
		if v == 'L' {
			current--
		} else {
			current++
		}
		ans = max(ans, current)
	}
	return ans

}
