package middle

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		temp := make([]int, 0)
		tempQueue := make([]*Node, 0)
		for i := 0; i < len(queue); i++ {
			temp = append(temp, queue[i].Val)
			tempQueue = append(tempQueue, queue[i].Children...)
		}
		result = append(result, temp)
		queue = tempQueue
	}
	return result

}
