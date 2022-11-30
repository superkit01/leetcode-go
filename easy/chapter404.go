package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	cache := make([]int, 0)
	sumOfLeftLeaves1(root, &cache, false)
	sum := 0
	for _, v := range cache {
		sum += v
	}
	return sum

}

func sumOfLeftLeaves1(node *TreeNode, cache *[]int, isLeft bool) {
	if node.Left == nil && node.Right == nil && isLeft {
		*cache = append(*cache, node.Val)
	}

	if node.Left != nil {
		sumOfLeftLeaves1(node.Left, cache, true)
	}
	if node.Right != nil {
		sumOfLeftLeaves1(node.Right, cache, false)
	}

}
