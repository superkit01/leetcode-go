package middle

/*
*
1000100
1000000
1111111
*/
func MaxActiveSectionsAfterTrade(s string) int {
	cnt1 := 0
	prev := -1
	max0 := 0
	curr := 0
	if s[0] == '0' {
		curr = 1
	} else {
		cnt1++
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			curr++
			if prev > 0 {
				max0 = max(max0, prev+curr)
			}
		} else {
			if s[i-1] == '0' {
				prev = curr
				curr = 0
			}
			cnt1++
		}

	}

	return max0 + cnt1

}
