package easy

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return recursive2(root, targetSum)

}

func recursive2(node *TreeNode, remaining int) bool {
	if node.Left == nil && node.Right == nil {
		if node.Val == remaining {
			return true
		} else {
			return false
		}
	}

	if node.Left != nil && node.Right == nil {
		return recursive2(node.Left, remaining-node.Val)
	}
	if node.Right != nil && node.Left == nil {
		return recursive2(node.Right, remaining-node.Val)
	} else {
		return recursive2(node.Left, remaining-node.Val) || recursive2(node.Right, remaining-node.Val)
	}

}
