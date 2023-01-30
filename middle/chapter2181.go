package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	var pre = head
	var current = head.Next

	for current.Next != nil {
		if current.Val != 0 {
			pre.Val = pre.Val + current.Val
		} else {
			pre.Next = current
			pre = pre.Next
		}

		current = current.Next
	}
	pre.Next = nil
	return head

}
