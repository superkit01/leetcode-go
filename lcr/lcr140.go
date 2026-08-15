package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainingPlanII(head *ListNode, cnt int) *ListNode {

	len := 0

	cur := head

	for cur != nil {
		len++
		cur = cur.Next
	}

	index := 0
	cur = head
	for cur != nil {
		if index == len-cnt {
			return cur
		}
		cur = cur.Next
		index++
	}

	return nil

}
