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
 *  前序  根 --> 左 --> 右
 */
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	recursive3(&result, root)
	return result
}

func recursive3(result *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	recursive3(result, root.Left)
	recursive3(result, root.Right)
}
