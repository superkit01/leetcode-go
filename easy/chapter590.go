package easy

func postorder(root *Node) []int {
	result := make([]int, 0)
	digui6(root, &result)
	return result
}

func digui6(root *Node, result *[]int) {
	if root == nil {
		return
	}

	for _, v := range root.Children {
		digui6(v, result)
	}

	*result = append(*result, root.Val)
}
