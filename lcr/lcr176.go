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
func isBalanced(root *TreeNode) bool {

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		ldeep := dfs(root.Left)
		rdeep := dfs(root.Right)
		if ldeep == -1 || rdeep == -1 || int(math.Abs(float64(ldeep-rdeep))) > 1 {
			return -1
		}
		return max(ldeep, rdeep) + 1
	}

	return dfs(root) != -1

}
