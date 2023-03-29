package middle

func countVowelStrings(n int) int {
	a, e, o, i, u := 1, 1, 1, 1, 1
	if n == 1 {
		return a + e + o + i + u
	}

	for v := 1; v < n; v++ {
		a = a + e + o + i + u
		e = e + o + i + u
		o = o + i + u
		i = i + u
	}
	return a + e + o + i + u
}
