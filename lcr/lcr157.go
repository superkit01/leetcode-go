package lcr

import (
	"sort"
)

func GoodsOrder(goods string) []string {
	n := len(goods)

	chars := []int{}
	for _, v := range goods {
		chars = append(chars, int(v))
	}
	sort.Ints(chars)

	ans := []string{}

	var dfs func(current string, bit int)
	dfs = func(current string, bit int) {
		if bit == 0 {
			ans = append(ans, current)
			return
		}

		pre := -1
		for i := 0; i < n; i++ {
			if pre != -1 && chars[i] == chars[pre] {
				continue
			}
			if bit>>i&0b1 == 0b0 {
				continue
			}

			dfs(current+string(rune(chars[i])), bit & ^(1<<i))
			pre = i
		}

	}

	dfs("", (1<<n)-1)

	return ans

}
