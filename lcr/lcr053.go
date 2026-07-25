package lcr

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode = nil
	var inorderDfs func(root *TreeNode, p *TreeNode) *TreeNode
	inorderDfs = func(root *TreeNode, p *TreeNode) *TreeNode {
		if root == nil {
			return ans
		}
		if root.Val > p.Val {
			if ans == nil || root.Val < ans.Val {
				ans = root
			}
			return inorderDfs(root.Left, p)
		} else {
			return inorderDfs(root.Right, p)
		}
	}
	inorderDfs(root, p)
	return ans
}
