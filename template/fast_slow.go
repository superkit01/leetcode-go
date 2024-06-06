package template

type ListNode struct {
	Val  int
	Next *ListNode
}

//  1 -> 2 -> 3 -> 4-> NULL

// https://blog.csdn.net/qq_54851255/article/details/122775138

// fast 先进入环 slow后进环 此时fast与slow之间的距离为 X ，之后每一步两者之间的距离 X 减1，最终一定相遇
// 判断链表是否有环
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 相遇时相遇点距离入环点X  C为环长
// slow ：L + X  fast： L + X + n*C
// fast是slow的两倍  L + X + n*C = 2 (L + X)  ==> (n-1)*C + C-X = L
// 找出链表中环的进入点
func findEntryOfCycle(head *ListNode) *ListNode {
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

	ans := head
	for ans != fast {
		ans = ans.Next
		fast = fast.Next
	}
	return ans
}

// 链表的中间节点
// 偶数节点时返回中间两节点的后一个
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
