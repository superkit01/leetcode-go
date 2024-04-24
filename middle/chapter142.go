package middle

func detectCycle(head *ListNode) *ListNode {

	fast := head
	slow := head

	hasCycle := false

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	for fast != head {
		fast = fast.Next
		head = head.Next
	}
	return fast

}
