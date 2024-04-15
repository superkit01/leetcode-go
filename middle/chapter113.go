package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	current := make([]int, 0)
	dfs9(root, targetSum, &ans, current)

	return ans

}

func dfs9(root *TreeNode, remain int, ans *[][]int, current []int) {

	if root.Left == nil && root.Right == nil {
		if remain != root.Val {
			return
		}
		current = append(current, root.Val)
		*ans = append(*ans, current)
		return

	}

	if root.Left != nil {
		tmp := make([]int, 0)
		tmp = append(tmp, current...)
		tmp = append(tmp, root.Val)
		dfs9(root.Left, remain-root.Val, ans, tmp)

	}

	if root.Right != nil {
		tmp := make([]int, 0)
		tmp = append(tmp, current...)
		tmp = append(tmp, root.Val)
		dfs9(root.Right, remain-root.Val, ans, tmp)
	}

}
