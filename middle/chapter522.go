package middle

import "sort"

func findLUSlength(strs []string) int {
	cnt := map[string]int{}

	distinct := []string{}

	for _, v := range strs {
		if cnt[v] == 0 {
			distinct = append(distinct, v)
		}
		cnt[v]++
	}

	sort.Slice(distinct, func(i, j int) bool {
		return len(distinct[i]) < len(distinct[j])
	})

	maxCount := len(distinct[len(distinct)-1])

	checkSubSeq := func(index int) bool {
		for i := len(distinct) - 1; i > index; i-- {
			if len(distinct[index]) == len(distinct[i]) {
				return false
			}

			//检查s是否是t的子序列
			check := func(s, t string) bool {
				m := 0 //s
				n := 0 //t
				for m < len(s) {
					if n == len(t) {
						return false
					}

					if s[m] == t[n] {
						m++
						n++
						continue
					}
					n++
				}
				return true
			}

			if check(distinct[index], distinct[i]) {
				return true
			}
		}
		return false

	}

	for i := len(distinct) - 1; i >= 0; i-- {
		if cnt[distinct[i]] > 1 {
			continue
		}

		if len(distinct[i]) == maxCount {
			return maxCount
		} else {
			if !checkSubSeq(i) {
				return len(distinct[i])
			}
		}
	}
	return -1

}
