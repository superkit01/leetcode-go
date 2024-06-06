package template

import "fmt"

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

//前序遍历  根-左-右 迭代法
func preOrderWithStack(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		fmt.Printf("%v ", pop.Val)
		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}
	}
}

//后序遍历  左-右-根  迭代法==> (前序微调【先压栈左子树在压栈右子树】 根-右-左   =>  在利用额外栈进行一次入栈出栈操作)
func postOrderWithStack(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	assistStack := make([]*TreeNode, 0)

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		//出栈之后入辅助栈
		assistStack = append(assistStack, pop)

		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}

		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
	}

	for len(assistStack) > 0 {
		pop := assistStack[len(stack)-1]
		fmt.Printf("%v ", pop.Val)
		assistStack = assistStack[0 : len(assistStack)-1]
	}
}

//中序遍历  左-根-右  迭代法
func inOrderWithStack(root *TreeNode) {
	// stack := make([]*TreeNode, 0)

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
