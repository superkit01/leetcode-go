package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func TrainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	curr := dummy

	head1 := l1
	head2 := l2

	for head1 != nil || head2 != nil {
		if head1 == nil {
			curr.Next = head2
			break
		}

		if head2 == nil {
			curr.Next = head1
			break
		}

		if head1.Val <= head2.Val {
			curr.Next = head1
			curr = curr.Next

			head1 = head1.Next
		} else {
			curr.Next = head2
			curr = curr.Next

			head2 = head2.Next
		}

	}

	return dummy.Next
}

func TrainningPlan_I(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {

		l1.Next = TrainningPlan_I(l1.Next, l2)
		return l1
	} else {
		l2.Next = TrainningPlan_I(l1, l2.Next)
		return l2
	}

}
