package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	//if root ==nil {
	//	return true
	//}
	//
	//if root.Left !=nil  && root.Left.Val>=root.Val {
	//	return false
	//}
	//
	//if root.Right !=nil  && root.Right.Val<=root.Val {
	//	return false
	//}
	//
	//return isValidBST(root.Left) && isValidBST(root.Right)

	if root == nil {
		return true
	}

	//中序遍历
	result := make([]int, 0)
	dfs2(root, &result)

	//升序验证
	temp := result[0]
	for i := 1; i < len(result); i++ {
		if result[i] <= temp {
			return false
		}
		temp = result[i]
	}
	return true

}

func dfs2(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	dfs2(node.Left, result)
	*result = append(*result, node.Val)
	dfs2(node.Right, result)
}

//func main() {
//	//[5,4,6,null,null,3,7]
//	isValidBST(&TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, &TreeNode{3, nil, nil}, &TreeNode{7, nil, nil}}})
//}
