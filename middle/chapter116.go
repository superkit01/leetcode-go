package middle

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	if root == nil {
		return root
	}

	queue := make([]*Node1, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tempQueue := make([]*Node1, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]

			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}

			if i <= len(queue)-2 {
				node.Next = queue[i+1]
			} else {
				node.Next = nil
			}
		}
		queue = tempQueue
	}

	return root

}
