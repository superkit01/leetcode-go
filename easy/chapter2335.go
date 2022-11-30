package easy

import "sort"

func fillCups(amount []int) int {
	sort.Ints(amount)

	sum := amount[0] + amount[1] + amount[2]

	if amount[0]+amount[1] >= amount[2] {
		return (sum + 1) / 2
	} else {
		return amount[2]
	}

}
