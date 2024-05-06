package lcr

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deduceTree(preorder []int, inorder []int) *TreeNode {
	return dfs(&preorder, inorder)
}

func dfs(preorder *[]int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	val := (*preorder)[0]
	root := &TreeNode{
		Val: val,
	}
	*preorder = (*preorder)[1:]

	for index, v := range inorder {
		if v == val {
			leftInorder := inorder[0:index]
			rightInorder := inorder[index+1:]
			if len(leftInorder) > 0 {
				root.Left = dfs(preorder, leftInorder)
			}
			if len(rightInorder) > 0 {
				root.Right = dfs(preorder, rightInorder)
			}
			break

		}
	}
	return root

}
