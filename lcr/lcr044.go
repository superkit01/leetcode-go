package lcr

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		max := math.MinInt32

		temp := []*TreeNode{}
		for _, v := range queue {
			if max < v.Val {
				max = v.Val
			}

			if v.Left != nil {
				temp = append(temp, v.Left)
			}

			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}

		ans = append(ans, max)

		queue = temp
	}

	return ans

}
