package lcr

func decorateRecordII(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := []*TreeNode{}
		tempAns := []int{}

		for _, v := range queue {
			tempAns = append(tempAns, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		ans = append(ans, tempAns)
		queue = temp

	}
	return ans
}
