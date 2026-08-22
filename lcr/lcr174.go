package lcr

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTargetNode(root *TreeNode, cnt int) int {
	ans := 0

	var midDfs func(root *TreeNode)
	midDfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		midDfs(root.Right)

		cnt--
		if cnt == 0 {
			ans = root.Val
		}
		midDfs(root.Left)

	}
	midDfs(root)

	return ans
}
