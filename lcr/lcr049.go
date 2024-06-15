package lcr

func sumNumbers(root *TreeNode) int {
	nums := []int{}

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			nums = append(nums, node.Val)
			return
		}
		if node.Left != nil {
			node.Left.Val = node.Val*10 + node.Left.Val
			dfs(node.Left)
		}

		if node.Right != nil {
			node.Right.Val = node.Val*10 + node.Right.Val
			dfs(node.Right)
		}
	}

	dfs(root)

	ans := 0
	for _, v := range nums {
		ans += v
	}
	return ans
}
