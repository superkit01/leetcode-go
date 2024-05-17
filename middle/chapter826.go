package middle

import (
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	diffProf := make([]*DiffProfit, 0)
	for i := range difficulty {
		diffProf = append(diffProf, &DiffProfit{difficulty[i], profit[i]})
	}

	sort.Slice(diffProf, func(i, j int) bool {
		if diffProf[i].difficulty == diffProf[j].difficulty {
			return diffProf[i].profit < diffProf[j].profit
		}
		return diffProf[i].difficulty < diffProf[j].difficulty
	})

	if len(diffProf) > 1 {
		for i := 0; i < len(diffProf)-1; i++ {
			if diffProf[i+1].profit < diffProf[i].profit {
				diffProf[i+1].profit = diffProf[i].profit
			}
		}
	}
	ans := 0
	for _, v := range worker {
		l := 0
		r := len(diffProf) - 1

		for l <= r {
			mid := (l + r) / 2
			if diffProf[mid].difficulty <= v {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if l != 0 {
			ans += diffProf[l-1].profit
		}
	}
	return ans

}

type DiffProfit struct {
	difficulty int
	profit     int
}
