package topics

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历  根-左-右
func preOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	preOrder(root.Left, ans)
	preOrder(root.Right, ans)
}

// 中序遍历  左-根-右
func inOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	preOrder(root.Left, ans)
	*ans = append(*ans, root.Val)
	preOrder(root.Right, ans)
}

// 后序遍历  左-右-跟
func postOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	preOrder(root.Left, ans)
	preOrder(root.Right, ans)
	*ans = append(*ans, root.Val)
}

// 层序遍历
func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}
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
		queue = temp
	}

}

// 从前序遍历和中序遍历恢复Tree
// 前序  根-左-右
// 中序  左-根-右
func reBuildTree(preOrder *[]int, inOrder []int) *TreeNode {
	if len(inOrder) == 0 {
		return nil
	}
	val := (*preOrder)[0]
	*preOrder = (*preOrder)[1:]

	root := &TreeNode{
		Val: val,
	}

	for i, v := range inOrder {
		if v == (*preOrder)[0] {
			root.Left = reBuildTree(preOrder, inOrder[0:i])
			root.Right = reBuildTree(preOrder, inOrder[i+1:])
			break
		}
	}
	return root
}
