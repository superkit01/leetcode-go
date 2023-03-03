package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddChain := &ListNode{}
	currentOdd := oddChain

	evenChain := &ListNode{}
	currentEven := evenChain

	current := head
	swipe := false

	for current != nil {
		if swipe {
			currentEven.Next = current
			currentEven = currentEven.Next

		} else {
			currentOdd.Next = current
			currentOdd = currentOdd.Next
		}
		swipe = !swipe
		current = current.Next
	}

	currentOdd.Next = evenChain.Next
	currentEven.Next = nil
	return oddChain.Next

}
