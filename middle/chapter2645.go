package middle

func addMinimum(word string) int {
	t := 0
	for i := 1; i < len(word); i++ {
		if word[i-1] >= word[i] {
			t++
		}
	}
	return 3*(t+1) - len(word)

}
