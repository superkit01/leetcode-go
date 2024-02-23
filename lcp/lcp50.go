package lcp

import "sort"

func giveGem(gem []int, operations [][]int) int {
	for _, v := range operations {
		transfer := gem[v[0]] / 2
		gem[v[0]] = gem[v[0]] - transfer
		gem[v[1]] = gem[v[1]] + transfer
	}

	sort.Ints(gem)
	return gem[len(gem)-1] - gem[0]

}
