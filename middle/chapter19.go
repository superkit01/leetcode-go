package middle

/**
* Definition for singly-linked list.
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := make(map[int]*ListNode)
	tmp, lastIndex := recursive(-1, head, tmp)
	if lastIndex+1 == n {
		return head.Next
	}

	preNode := tmp[lastIndex-n]

	if _, ok := tmp[lastIndex-(n-2)]; ok {
		afterNode := tmp[lastIndex-(n-2)]
		preNode.Next = afterNode
	} else {
		preNode.Next = nil
	}
	return head
}

func recursive(index int, listNode *ListNode, tmp map[int]*ListNode) (map[int]*ListNode, int) {
	if listNode != nil {
		index = index + 1
		tmp[index] = listNode

		tmp, index = recursive(index, listNode.Next, tmp)
		return tmp, index
	} else {
		return tmp, index
	}

}
