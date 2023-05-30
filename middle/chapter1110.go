package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

	result := make([]*TreeNode, 0)

	dfs3(nil, root, &result, to_delete)

	if !contains(to_delete, root.Val) {
		result = append(result, root)
	}

	return result

}

func dfs3(parent *TreeNode, root *TreeNode, result *[]*TreeNode, to_delete []int) {

	if root.Left != nil {
		dfs3(root, root.Left, result, to_delete)
	}

	if root.Right != nil {
		dfs3(root, root.Right, result, to_delete)
	}

	if contains(to_delete, root.Val) {
		if parent != nil && parent.Left != nil && parent.Left == root {
			parent.Left = nil
		}
		if parent != nil && parent.Right != nil && parent.Right == root {
			parent.Right = nil
		}
		if root.Left != nil {
			*result = append(*result, root.Left)
		}
		if root.Right != nil {
			*result = append(*result, root.Right)
		}

	}

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
