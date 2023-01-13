package easy

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := make([]int, 0)
	digui5(root, &result)
	return result

}

func digui5(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	for _, v := range node.Children {
		digui5(v, result)
	}

}
