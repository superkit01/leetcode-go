package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	container := make(map[*ListNode]interface{})

	iterA := headA
	for {
		if iterA == nil {
			break
		}
		container[iterA] = nil
		iterA = iterA.Next
	}

	iterB := headB
	for {
		if iterB == nil {
			break
		}

		if _, ok := container[iterB]; ok {
			return iterB
		}

		iterB = iterB.Next
	}

	return nil

}
