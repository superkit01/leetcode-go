package lcr

func decorateRecordIII(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}

	direction := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := []*TreeNode{}
		tempAns := []int{}

		for _, v := range queue {
			if direction > 0 {
				tempAns = append(tempAns, v.Val)
			} else {
				tempAns = append([]int{v.Val}, tempAns...)
			}

			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		ans = append(ans, tempAns)
		queue = temp
		direction *= -1

	}
	return ans
}
