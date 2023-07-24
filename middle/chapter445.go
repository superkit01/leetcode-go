package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1 = ReverseList(l1)
	l2 = ReverseList(l2)

	currentl1 := l1
	currentl2 := l2
	jin := 0

	result := &ListNode{Val: 0, Next: nil}
	head := result

	for currentl1 != nil || currentl2 != nil || jin != 0 {
		val := jin
		jin = 0
		if currentl1 != nil {
			val += currentl1.Val
		}
		if currentl2 != nil {
			val += currentl2.Val
		}
		if val >= 10 {
			jin = val / 10
			val = val % 10
		}

		head.Next = &ListNode{Val: val, Next: nil}
		head = head.Next

		if currentl1 != nil {
			currentl1 = currentl1.Next
		}

		if currentl2 != nil {
			currentl2 = currentl2.Next
		}
	}
	result = result.Next

	return ReverseList(result)

}

func ReverseList(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	tail := node
	current := node.Next

	tail.Next = nil

	for current != nil {
		next := current.Next
		current.Next = tail
		tail = current
		current = next
	}
	return tail
}
