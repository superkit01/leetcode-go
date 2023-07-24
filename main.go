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

	easy.ArithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3)

	middle.IsRobotBounded("GGLLGG")

	middle.NumberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}, {4, 4}})

	easy.MostFrequentEven([]int{0, 1, 2, 2, 4, 4, 1})
	easy.SortPeople([]string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"},
		[]int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224})

	middle.ReverseList(&middle.ListNode{Val: 1, Next: &middle.ListNode{Val: 2, Next: &middle.ListNode{Val: 3, Next: &middle.ListNode{Val: 4, Next: nil}}}})
}
