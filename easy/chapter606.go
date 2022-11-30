package easy

import (
	"strconv"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

//      1
//     /   \
//    2     3
//   /	\
//  4	 4
//
// 输出: "1(2(4))(3)"

func tree2str(root *TreeNode) string {
	return dfs(root)
}

func dfs(node *TreeNode) string {
	if node == nil {
		return ""
	}
	s := strconv.Itoa(node.Val)
	if node.Left != nil || node.Right != nil {
		s += "(" + dfs(node.Left) + ")"
	}
	if node.Right != nil {
		s += "(" + dfs(node.Right) + ")"
	}
	return s
}

//func main() {
//	root := TreeNode{Val: 1,
//		Left: &TreeNode{Val: 2,
//			Left: &TreeNode{Val: 4}},
//		Right: &TreeNode{Val: 3}}
//
//	fmt.Printf("%v \n", tree2str(&root))
//}
