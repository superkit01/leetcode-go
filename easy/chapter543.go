package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lDeep := dfs(root.Left)
		rDeep := dfs(root.Right)
		ans = max(ans, lDeep+rDeep)

		return max(lDeep, rDeep) + 1
	}

	dfs(root)

	return ans
}

