package easy

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	iter := head.Next
	prev := head
	for iter != nil {
		if iter.Val != val {
			prev = iter
			iter = iter.Next
		} else {
			iter = iter.Next
			prev.Next = iter
		}

	}

	if head.Val == val {
		return head.Next
	}

	return head

}
