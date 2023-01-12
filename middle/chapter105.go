package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	return digui(&preorder, inorder)
}

func digui(preorder *[]int, inorder []int) *TreeNode {
	root := (*preorder)[0]
	*preorder = (*preorder)[1:]
	tree := TreeNode{
		Val: root,
	}

	index := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root {
			index = i
		}
	}

	if index > 0 {
		left := inorder[0:index]
		tree.Left = digui(preorder, left)

	}

	if index < len(inorder)-1 {
		right := inorder[index+1:]
		tree.Right = digui(preorder, right)
	}
	return &tree
}
