package template

//s1是否存在s2的子串
func KMP(s1 string, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	}
	//构建next数组
	next := buildNext(s2)

	pos1 := 0
	pos2 := 0
	for {
		if pos2 >= len(s2) {
			return true
		}

		if pos1 >= len(s1) {
			return false
		}

		if s1[pos1] == s2[pos2] {
			pos1++
			pos2++
			continue
		}

		if pos2 == 0 {
			pos1++
			continue
		}
		pos2 = next[pos2]
	}
}

// Next[index]:不含当前位置,不含整体 的最长公共前后缀  0~index-1
func buildNext(s string) []int {
	if len(s) == 1 {
		return []int{-1}
	}

	next := make([]int, len(s))
	next[0] = -1
	next[1] = 0

	pos := 0
	i := 2
	for i < len(s) {
		if s[i-1] == s[pos] {
			next[i] = pos + 1
			pos++
			i++
			continue
		}

		if pos == 0 {
			next[i] = 0
			i++
			continue
		}

		if pos > 0 {
			pos = next[pos]
		}
	}

	return next

}
