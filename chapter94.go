package main

/**
 * Definition for a binary tree node.

 */
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/**
	前序  根--左--右
	中序  左--根--右
	后序  左--右--根
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	digui(root, &result)
	return result
}

func digui(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	digui(root.Left, result)
	*result = append(*result, root.Val)
	digui(root.Right, result)
}

//
//func main(){
//	root:=TreeNode{1,&TreeNode{2,&TreeNode{3,nil,nil},&TreeNode{4,nil,nil}},&TreeNode{2,nil,nil}}
//	result:=inorderTraversal(&root)
//	fmt.Println(result)
//}