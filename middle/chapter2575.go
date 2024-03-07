package middle

func divisibilityArray(word string, m int) []int {
	ans := make([]int, 0)
	current := 0
	for _, v := range word {
		current = current*10 + int(v-'0')
		current = current % m

		if current == 0 {
			ans = append(ans, 1)
		} else {
			ans = append(ans, 0)
		}
	}
	return ans

}
