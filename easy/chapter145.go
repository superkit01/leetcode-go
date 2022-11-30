package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 *  后序   左 --> 右--> 根
 */
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	recursive4(&result, root)
	return result
}

func recursive4(result *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	recursive4(result, root.Left)
	recursive4(result, root.Right)
	*result = append(*result, root.Val)
}
