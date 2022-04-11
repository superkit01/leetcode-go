package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	i := 0
outer:
	for len(queue) > 0 {
		i++

		tempQueue := make([]*TreeNode, 0)

		for _, e := range queue {
			if e.Left == nil && e.Right == nil {
				break outer
			}
			if e.Left != nil {
				tempQueue = append(tempQueue, e.Left)
			}
			if e.Right != nil {
				tempQueue = append(tempQueue, e.Right)
			}
		}

		queue = tempQueue
	}
	return i
}
