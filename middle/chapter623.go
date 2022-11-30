package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}

	level := 2

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		if level == depth {
			for i := 0; i < len(queue); i++ {
				temp := queue[i]
				tempLeft := temp.Left
				temp.Left = &TreeNode{
					Val:   val,
					Left:  tempLeft,
					Right: nil,
				}
				tempRight := temp.Right
				temp.Right = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: tempRight,
				}
			}
			return root
		} else {
			tempQueue := make([]*TreeNode, 0)
			for i := 0; i < len(queue); i++ {
				temp := queue[i]
				if temp.Left != nil {
					tempQueue = append(tempQueue, temp.Left)
				}
				if temp.Right != nil {
					tempQueue = append(tempQueue, temp.Right)
				}
			}
			queue = make([]*TreeNode, 0)
			queue = append(queue, tempQueue...)
		}
		level++
	}
	return root

}

// func main() {
// addOneRow(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, 5, 4)
// }
