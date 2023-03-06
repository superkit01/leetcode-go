package week335

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	sum := make([]int64, 0)
	sum = append(sum, int64(root.Val))

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		temp := make([]*TreeNode, 0)

		for _, v := range queue {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}

		if len(temp) > 0 {
			var tempSum int64
			tempSum = 0
			for _, e := range temp {
				tempSum += int64(e.Val)
			}
			sum = append(sum, int64(tempSum))
		}
		queue = temp
	}

	if k > len(sum) {
		return -1
	} else {
		sort.Slice(sum, func(i, j int) bool { return sum[i] > sum[j] })
		return sum[k-1]
	}

}
