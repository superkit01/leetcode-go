package middle

func connect1(root *Node1) *Node1 {
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

func connect2(root *Node1) *Node1 {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	queue := make([]*Node1, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		tempQueue := make([]*Node1, 0)
		for _, v := range queue {
			if v.Left != nil {
				tempQueue = append(tempQueue, v.Left)
			}
			if v.Right != nil {
				tempQueue = append(tempQueue, v.Right)
			}
		}

		if len(tempQueue) > 1 {
			for i := 0; i < len(tempQueue)-1; i++ {
				if tempQueue[i] != nil {
					tempQueue[i].Next = tempQueue[i+1]
				}
			}
		}
		queue = tempQueue
	}

	return root

}
