package easy

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	temp := strconv.Itoa(root.Val)

	recursive5(root, &result, temp)
	return result
}

func recursive5(node *TreeNode, result *[]string, temp string) {
	if node.Left == nil && node.Right == nil {
		(*result) = append((*result), temp)
		return
	}
	if node.Left != nil {
		recursive5(node.Left, result, temp+"->"+strconv.Itoa(node.Left.Val))
	}
	if node.Right != nil {
		recursive5(node.Right, result, temp+"->"+strconv.Itoa(node.Right.Val))
	}
}
