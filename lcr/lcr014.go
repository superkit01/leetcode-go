package lcr

// 滑动窗口
func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	count1 := [26]int{}
	count2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		if i > 0 {
			count2[s2[i-1]-'a']--
			count2[s2[i+len(s1)-1]-'a']++
		}

		if count1 == count2 {
			return true
		}
	}

	return false
}
