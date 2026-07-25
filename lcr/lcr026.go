package lcr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	nodes := make([]*ListNode, 0)

	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	i, j := 0, len(nodes)-1

	for i < j {
		nodes[i].Next = nodes[j]
		i++

		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
