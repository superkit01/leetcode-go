package easy

// bfs
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	length := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		length++
		tempQueue := make([]*TreeNode, 0)

		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				tempQueue = append(tempQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tempQueue = append(tempQueue, queue[i].Right)
			}
		}
		queue = tempQueue
	}

	return length

}

// dfs
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1

}
