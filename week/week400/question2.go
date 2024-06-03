package week400

import "sort"

func countDays(days int, meetings [][]int) int {

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})

	ans := 0

	i := 1

	for _, v := range meetings {
		if v[0] > i {
			ans += v[0] - i

		}
		i = max(i, v[1]+1)
	}

	if i <= days {
		ans += days - i + 1
	}
	return ans

}
