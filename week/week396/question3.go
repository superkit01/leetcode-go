package week396

func minAnagramLength(s string) int {
	n := len(s)
	prim := make([]int, 0)
	prim = append(prim, 1)
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			prim = append(prim, i)
		}
	}

outer:
	for i := 0; i < len(prim); i++ {
		k := prim[i]
		for j := 0; j < n/k-1; j++ {
			if !equals(s[j*k:j*k+k], s[(j+1)*k:(j+1)*k+k]) {
				continue outer
			}
		}
		return k
	}

	return len(s)

}

func equals(s1, s2 string) bool {
	count1 := make(map[rune]int, 0)
	count2 := make(map[rune]int, 0)
	for _, v := range s1 {
		count1[v]++
	}
	for _, v := range s2 {
		count2[v]++
	}
	for k, v := range count1 {
		if count2[k] != v {
			return false
		}
	}
	return true
}
