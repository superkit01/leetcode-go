package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome2(head *ListNode) bool {
	list := make([]*ListNode, 0)

	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	for i := 0; i < len(list)/2; i++ {
		if list[i].Val != list[len(list)-1-i].Val {
			return false
		}
	}
	return true
}
