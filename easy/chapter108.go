package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

	return buildTree(nums, 0, len(nums))
}

func buildTree(nums []int, L int, R int) *TreeNode {
	if R-L <= 1 {
		return &TreeNode{nums[L], nil, nil}
	}

	mid := L + (R-L-1)/2
	root := &TreeNode{nums[mid], nil, nil}
	if mid-L > 0 {
		root.Left = buildTree(nums, L, mid)
	}
	if R-mid > 0 {
		root.Right = buildTree(nums, mid+1, R)
	}
	return root
}

//func main() {
//	root := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5})
//	fmt.Println("%v", root)
//}
