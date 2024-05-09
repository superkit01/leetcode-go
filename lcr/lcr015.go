package lcr

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	ans := []int{}

	countS := [26]int{}
	countP := [26]int{}

	for i := 0; i < len(p); i++ {
		countS[s[i]-'a']++
		countP[p[i]-'a']++
	}

	for i := 0; i <= len(s)-len(p); i++ {
		if i > 0 {
			countS[s[i-1]-'a']--
			countS[s[i+len(p)-1]-'a']++
		}
		if countS == countP {
			ans = append(ans, i)
		}

	}
	return ans

}
