package middle

import "sort"

func findWinners(matches [][]int) [][]int {
	cnt := make(map[int]int, 0)
	for _, v := range matches {
		if _, ok := cnt[v[0]]; !ok {
			cnt[v[0]] = 0
		}
		cnt[v[1]]++
	}

	ans := [][]int{{}, {}}

	for k, v := range cnt {
		if v == 1 {
			ans[1] = append(ans[1], k)
		}
		if v == 0 {
			ans[0] = append(ans[0], k)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
