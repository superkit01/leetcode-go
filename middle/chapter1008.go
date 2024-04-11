package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{
		Val: preorder[0],
	}

	buildTree1(preorder[1:], root)

	return root
}

func buildTree1(preorder []int, root *TreeNode) {
	if len(preorder) == 0 {
		return
	}

	index := -1
	for i := range preorder {
		if preorder[i] > root.Val {
			index = i
			break
		}
	}

	if index == -1 {
		root.Left = &TreeNode{
			Val: preorder[0],
		}
		if len(preorder) > 1 {
			buildTree1(preorder[1:], root.Left)
		}

	} else if index == 0 {
		root.Right = &TreeNode{
			Val: preorder[0],
		}
		if len(preorder) > 1 {
			buildTree1(preorder[1:], root.Right)
		}

	} else {
		left := preorder[0:index]
		root.Left = &TreeNode{
			Val: left[0],
		}
		if len(left) > 1 {
			buildTree1(left[1:], root.Left)
		}

		right := preorder[index:]
		root.Right = &TreeNode{
			Val: right[0],
		}
		if len(right) > 1 {
			buildTree1(right[1:], root.Right)
		}
	}
}
