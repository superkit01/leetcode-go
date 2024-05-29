package middle

func MaximumLength(s string) int {
	l := 1
	r := len(s) - 1

	for l < r {
		mid := (l + r) / 2
		if check(s, mid) {
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

func check(s string, l int) bool {
	cnt := make(map[byte]int)
outer:
	for i := 0; i <= len(s)-l; i++ {
		for j := i; j < l+i; j++ {
			if s[i] != s[j] {
				continue outer
			}
		}
		cnt[s[i]]++
		if cnt[s[i]] == 3 {
			return true
		}
	}

	return false
}
