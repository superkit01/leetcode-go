package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainningPlanIII(head *ListNode) *ListNode {

	var reverseList func(head *ListNode) *ListNode
	reverseList = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		newHead := reverseList(head.Next)

		tail := head.Next
		tail.Next = head

		head.Next = nil

		return newHead
	}

	return reverseList(head)

}
