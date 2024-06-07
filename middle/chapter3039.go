package middle

func lastNonEmptyString(s string) string {
	cnt := [26]int{}
	max := 0
	for _, v := range s {
		cnt[v-'a']++
		if max < cnt[v-'a'] {
			max = cnt[v-'a']
		}
	}

	cnt1 := make(map[int]int, 0)

	for i, v := range cnt {
		if v == max {
			cnt1[i] = max
		}
	}
	ans := ""
	for _, v := range s {
		if _, ok := cnt1[int(v-'a')]; ok {
			cnt1[int(v-'a')]--
			if cnt1[int(v-'a')] == 0 {
				ans += string(v)
			}
		}
	}
	return ans

}
