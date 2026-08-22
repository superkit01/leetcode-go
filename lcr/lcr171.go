package lcr

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA := headA
	pointB := headB

	for pointA != pointB {
		if pointA == nil {
			pointA = headB
		} else {
			pointA = pointA.Next
		}
		if pointB == nil {
			pointB = headA
		} else {
			pointB = pointB.Next
		}

	}
	return pointA

}
