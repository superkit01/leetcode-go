package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)

	return head
}

func deleteNodeII(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}

	return head
}
