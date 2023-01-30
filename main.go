package main

import (
	"leecode-go/middle"
)

func main() {
	// middle.MinMoves2([]int{1, 10, 2, 9})
	middle.MergeInBetween(&middle.ListNode{0, &middle.ListNode{1, &middle.ListNode{2, nil}}},
		1, 1, &middle.ListNode{1000000, &middle.ListNode{1000001, &middle.ListNode{1000002, &middle.ListNode{1000003, nil}}}})
}

/**
[0,1,2]
1
1
[1000000,1000001,1000002,1000003]
*/
