package middle

func maximumLength(s string) int {
	//各字符连续出现次数
	cnt := make(map[byte][]int, 0)
	k := 0
	for i := 0; i < len(s); i++ {
		if s[i] == s[k] {
			continue
		}
		cnt[s[k]] = append(cnt[s[k]], i-k)
		k = i
	}
	cnt[s[k]] = append(cnt[s[k]], len(s)-k)


	l := 1
	r := len(s) - 1

	for l < r {
		mid := (l + r) / 2
		if checkCnt(mid, cnt) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if l-1 == 0 {
		return -1
	}
	return l - 1

}

func checkCnt(mid int, cnt map[byte][]int) bool {
	for _, v := range cnt {
		c := 0
		for _, m := range v {
			if m >= mid+2 {
				return true
			}

			if m == mid+1 {
				c += 2
			}
			if m == mid {
				c++
			}
		}
		if c >= 3 {
			return true
		}
	}

	return false

}
