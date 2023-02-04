package middle

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	max := 0

	for _, v := range coins {
		if v <= max+1 {
			max = max + v
		} else {
			break
		}

	}
	return max + 1

}
