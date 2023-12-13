package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func BstToGst(root *TreeNode) *TreeNode {

	var sum = 0

	dfs8(root, &sum)
	return root

}

func dfs8(root *TreeNode, sum *int) {
	if root.Right != nil {
		dfs8(root.Right, sum)
	}
	*sum += root.Val
	root.Val = *sum

	if root.Left != nil {
		dfs8(root.Left, sum)
	}
}
