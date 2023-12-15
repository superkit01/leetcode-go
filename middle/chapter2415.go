package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func ReverseOddLevels(root *TreeNode) *TreeNode {

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	level := 1

	for len(queue) > 0 {
		tempQueue := make([]*TreeNode, 0)

		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				tempQueue = append(tempQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tempQueue = append(tempQueue, queue[i].Right)
			}

		}

		if level%2 == 1 {
			for j := 0; j < len(tempQueue)/2; j++ {
				tempQueue[j].Val, tempQueue[len(tempQueue)-1-j].Val = tempQueue[len(tempQueue)-1-j].Val, tempQueue[j].Val
			}
		}
		level++

		queue = tempQueue
	}

	return root

}
