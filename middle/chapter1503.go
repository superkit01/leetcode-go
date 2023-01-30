package middle

func getLastMoment(n int, left []int, right []int) int {
	t := -1
	for _, v := range left {
		if v > t {
			t = v
		}
	}

	for _, v := range right {
		if n-v > t {
			t = n - v
		}
	}
	return t

}
