package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	if root == nil {
		return 0
	}

	if root.Val > high {
		sum += rangeSumBST(root.Left, low, high)
	} else if root.Val < low {
		sum += rangeSumBST(root.Right, low, high)
	} else {
		sum += root.Val
		sum += rangeSumBST(root.Left, low, high)
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum

}
