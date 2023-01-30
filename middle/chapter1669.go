package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	index := 0

	var aNode *ListNode
	var bNode *ListNode
	current := list1
	for {
		index += 1
		if index == a {
			aNode = current
		}
		if index == b {
			bNode = current
			break
		}
		current = current.Next
	}

	last := list2
	for last.Next != nil {
		last = last.Next
	}
	last.Next = bNode.Next.Next
	aNode.Next = list2

	return list1

}
