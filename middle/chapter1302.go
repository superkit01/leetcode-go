package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	sum := 0
	for len(queue) != 0 {
		sum = 0
		tempQueue := make([]*TreeNode, 0)
		for len(queue) > 0 {
			temp := queue[0]
			queue = queue[1:]
			sum += temp.Val
			if temp.Left != nil {
				tempQueue = append(tempQueue, temp.Left)
			}
			if temp.Right != nil {
				tempQueue = append(tempQueue, temp.Right)
			}
		}
		queue = tempQueue
	}
	return sum

}
