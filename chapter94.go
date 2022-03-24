package main

import "fmt"

/**
 * Definition for a binary tree node.

 */
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	digui(root, &result)
	return result
}

func digui(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	digui(root.Left, result)
	digui(root.Right, result)
}


func main(){
	root:=TreeNode{1,&TreeNode{2,&TreeNode{3,nil,nil},&TreeNode{4,nil,nil}},&TreeNode{2,nil,nil}}
	result:=inorderTraversal(&root)
	fmt.Println(result)
}