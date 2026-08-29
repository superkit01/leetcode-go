package lcr

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func calculateDepth(root *TreeNode) int {
	ans := 0

	var dfs func(root *TreeNode, deep int)
	dfs = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		deep++
		ans = max(deep, ans)

		dfs(root.Left, deep)
		dfs(root.Right, deep)

	}

	dfs(root, 0)

	return ans
}
