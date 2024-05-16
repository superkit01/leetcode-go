package lcr

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	nums []*TreeNode
}

func Constructor043(root *TreeNode) CBTInserter {

	//BFS
	nums := []*TreeNode{}
	nums = append(nums, root)
	i := 0
	for i < len(nums) {
		if nums[i].Left != nil {
			nums = append(nums, nums[i].Left)
		}
		if nums[i].Right != nil {
			nums = append(nums, nums[i].Right)
		}
		i++
	}

	return CBTInserter{
		nums: nums,
	}

}

func (this *CBTInserter) Insert(v int) int {
	node := &TreeNode{Val: v}

	this.nums = append(this.nums, node)

	parentNode := this.nums[len(this.nums)/2-1]
	if len(this.nums)%2 == 0 {
		parentNode.Left = node
	} else {
		parentNode.Right = node
	}

	return parentNode.Val

}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.nums[0]
}
