package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	if root.Val != 0 && root.Val != 1 {
		if root.Val == 2 {
			return evaluateTree(root.Left) || evaluateTree(root.Right)
		} else {
			return evaluateTree(root.Left) && evaluateTree(root.Right)
		}
	} else {
		if root.Val == 0 {
			return false
		} else {
			return true
		}
	}
}
