package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		temp := make([]int, 0)
		tempQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				tempQueue = append(tempQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tempQueue = append(tempQueue, queue[i].Right)
			}
		}
		result = append(result, temp)
		queue = tempQueue
	}

	return result

}
