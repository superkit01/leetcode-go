package lcr

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	hasCyle := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCyle = true
			break
		}
	}
	if !hasCyle {
		return nil
	}

	ans := head

	for ans != fast {
		ans = ans.Next
		fast = fast.Next
	}
	return ans

}
