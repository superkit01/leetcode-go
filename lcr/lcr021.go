package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	deque := make([]*ListNode, 0)

	curr := head
	for curr != nil {
		deque = append(deque, curr)
		curr = curr.Next
	}

	if n == len(deque) {
		return deque[1]
	}

	if n == 1 {
		deque[len(deque)-2].Next = nil
		return head
	} else {
		deque[len(deque)-1-n].Next = deque[len(deque)-n].Next
		return head
	}

}
