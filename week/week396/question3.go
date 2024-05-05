package week396

func minAnagramLength(s string) int {
	n := len(s)
	prim := make([]int, 0)
	// for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
	// 	if n%i == 0 {
	// 		prim = append(prim, i)
	// 		prim = append(prim, n/i)
	// 	}
	// }
	// sort.Ints(prim)

	for i := 1; i <= n-1; i++ {
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

	count1 := [26]int{}
	count2 := [26]int{}

	for _, v := range s1 {
		count1[v-'a']++
	}
	for _, v := range s2 {
		count2[v-'a']++
	}
	return count1 == count2

	// count1 := make(map[rune]int, 0)
	// count2 := make(map[rune]int, 0)
	// for _, v := range s1 {
	// 	count1[v]++
	// }
	// for _, v := range s2 {
	// 	count2[v]++
	// }
	// for k, v := range count1 {
	// 	if count2[k] != v {
	// 		return false
	// 	}
	// }
	// return true
}
