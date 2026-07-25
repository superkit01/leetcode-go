package lcr

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	ans := math.MinInt
	maxDeep := 0

	var dfs func(node *TreeNode, deep int)

	dfs = func(node *TreeNode, deep int) {
		if node == nil {
			return
		}
		deep++
		dfs(node.Left, deep)
		dfs(node.Right, deep)
		if deep > maxDeep {
			maxDeep = deep
			ans = node.Val
		}
	}

	dfs(root, 0)
	return ans
}
