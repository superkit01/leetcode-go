package lcr

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if B == nil {
		return false
	}
	if A == nil {
		return false
	}

	var dfs func(A *TreeNode, B *TreeNode) bool
	dfs = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil {
			return false
		}

		if A.Val != B.Val {
			return false
		}
		return dfs(A.Left, B.Left) || dfs(A.Right, B.Right)
	}

	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
