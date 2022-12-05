package middle

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	direct := 1

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		tempQueue := make([]*TreeNode, 0)
		tempResult := make([]int, len(queue))
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
			if direct == 1 {
				tempResult[i] = node.Val
			} else {
				tempResult[len(tempResult)-1-i] = node.Val
			}
		}
		direct *= -1
		queue = tempQueue
		result = append(result, tempResult)
	}

	return result
}
