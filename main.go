package main

import (
	"fmt"
	"leetcode-go/easy"
	"leetcode-go/middle"
	"leetcode-go/week/week336"
)

func main() {

	week336.BeautifulSubarrays([]int{4, 3, 1, 2, 4})

	fmt.Print(easy.MinNumberOfHours(5, 1, []int{1, 3, 3}, []int{1, 3, 7}))

	middle.MaximalNetworkRank(4,
		[][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}})
}
