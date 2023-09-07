package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, result := dfs5(root)
	return result
}

func dfs5(node *TreeNode) (int, *TreeNode) {

	if node == nil {
		return 0, nil
	}
	left, leftNode := dfs5(node.Left)
	right, rightNode := dfs5(node.Right)
	if left < right {
		return right + 1, rightNode
	}
	if left > right {
		return left + 1, leftNode
	}
	return left + 1, node
}
