package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		temp := make([]*TreeNode, 0)
		tempResult := make([]int, 0)
		for _, v := range queue {
			tempResult = append(tempResult, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)

			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}

		queue = temp
		result = append(result, tempResult)

	}

	return reverse1(result)

}

func reverse1(result [][]int) [][]int {
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}

// func main() {
// 	root := &TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val: 2,
// 			Left: &TreeNode{
// 				Val: 4,
// 			},
// 			Right: &TreeNode{
// 				Val: 5,
// 			},
// 		},
// 		Right: &TreeNode{
// 			Val: 3,
// 			Left: &TreeNode{
// 				Val: 6,
// 			},
// 			Right: &TreeNode{
// 				Val: 7,
// 			},
// 		},
// 	}

// 	root = nil
// 	fmt.Printf("%v", levelOrderBottom(root))
// }
