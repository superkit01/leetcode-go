package lcr

func decorateRecord(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := make([]*TreeNode, 0)
		for _, v := range queue {
			ans = append(ans, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		queue = temp
	}
	return ans
}
