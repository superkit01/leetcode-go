package easy

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)

	i := 0
	j := 0
	result := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			result++
			i++
			j++
		} else {
			j++
		}
	}
	return result

}
