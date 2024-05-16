package lcr

func PathTarget(root *TreeNode, target int) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	var dfs func(node *TreeNode, path []int, rest int)

	dfs = func(node *TreeNode, path []int, rest int) {
		path = append(path, node.Val)
		rest -= node.Val
		if node.Left == nil && node.Right == nil {
			if rest == 0 {
				temp := []int{}
				temp = append(temp, path...)
				ans = append(ans, temp)
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, path, rest)
		}

		if node.Right != nil {
			dfs(node.Right, path, rest)
		}

	}

	dfs(root, []int{}, target)

	return ans
}
