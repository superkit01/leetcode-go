package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBookList(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	if head.Next == nil {
		return []int{head.Val}
	}

	var reverseLink func(head *ListNode) *ListNode

	reverseLink = func(head *ListNode) *ListNode {
		if head.Next == nil {
			return head
		}

		newHead := reverseLink(head.Next)

		head.Next.Next = head
		head.Next = nil
		return newHead

	}

	reversedHead := reverseLink(head)

	ans := make([]int, 0)
	for reversedHead != nil {
		ans = append(ans, reversedHead.Val)
		reversedHead = reversedHead.Next
	}
	return ans
}
