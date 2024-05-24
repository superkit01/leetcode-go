package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//递归
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	next := head.Next
	next.Next = head
	head.Next = nil

	return newHead

}

//迭代
func reverseListII(head *ListNode) *ListNode {

	var pre *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre

}
