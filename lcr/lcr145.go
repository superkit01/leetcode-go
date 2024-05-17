package lcr

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkSymmetricTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var checkMirror func(node1, node2 *TreeNode) bool

	checkMirror = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		return checkMirror(node1.Left, node2.Right) && checkMirror(node1.Right, node2.Left)
	}

	return checkMirror(root.Left, root.Right)

}
