package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	return digui2(inorder, &postorder)
}

func digui2(inorder []int, postorder *[]int) *TreeNode {
	root := (*postorder)[len(*postorder)-1]
	*postorder = (*postorder)[0 : len(*postorder)-1]

	tree := TreeNode{
		Val: root,
	}

	index := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root {
			index = i
			break
		}
	}

	if index < len(inorder)-1 {
		right := inorder[index+1:]
		tree.Right = digui2(right, postorder)
	}

	if index > 0 {
		left := inorder[0:index]
		tree.Left = digui2(left, postorder)
	}

	return &tree
}
