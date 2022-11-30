package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	node  *TreeNode
	queue []*TreeNode
	cur   int
}

func Constructor4(root *TreeNode) CBTInserter {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	cur := 0
	index := 0
	for index < len(queue) {
		current := queue[index]
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
		index++
		if current.Left != nil && current.Right != nil {
			cur++
		}
	}

	return CBTInserter{
		node:  root,
		queue: queue,
		cur:   cur,
	}
}

func (this *CBTInserter) Insert(val int) int {
	newNode := &TreeNode{
		Val: val,
	}

	node := this.queue[this.cur]

	if node.Left == nil {
		node.Left = newNode
	} else {
		node.Right = newNode
		this.cur++
	}

	this.queue = append(this.queue, newNode)

	return node.Val

}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.queue[0]

}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
