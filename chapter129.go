package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	temp := make([]int, 0)

	dfs4(root, 0, &temp)

	sum := 0
	for _, i := range temp {
		sum += i
	}
	return sum

}

func dfs4(node *TreeNode, val int, temp *[]int) {
	if node.Left == nil && node.Right == nil {
		*temp = append(*temp, val*10+node.Val)
		return
	}

	if node.Left != nil {
		dfs4(node.Left, val*10+node.Val, temp)
	}
	if node.Right != nil {
		dfs4(node.Right, val*10+node.Val, temp)
	}

}
