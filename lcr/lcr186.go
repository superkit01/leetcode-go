package lcr

func checkDynasty(places []int) bool {
	m := make(map[int]int)
	mx := -1
	mi := 100
	for _, v := range places {
		if v == 0 {
			continue
		}
		mx = max(mx, v)
		mi = min(mi, v)
		if _, ok := m[v]; ok {
			return false
		}
		m[v]++
	}

	return mx-mi < 5

}
