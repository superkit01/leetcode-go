package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DistanceK(root *TreeNode, target *TreeNode, k int) []int {
	container := make(map[int][]int, 0)
	if root != nil {
		dfs6(root, root.Left, container)
		dfs6(root, root.Right, container)
	}

	if _, ok := container[target.Val]; !ok {
		return []int{}
	}

	result := make([]int, 0)

	dfs7(-1, target.Val, container, 0, k, &result)

	return result

}

func dfs6(pNode *TreeNode, node *TreeNode, container map[int][]int) {
	if node == nil {
		return
	}

	if _, ok := container[node.Val]; !ok {
		container[node.Val] = make([]int, 0)

	}
	container[node.Val] = append(container[node.Val], pNode.Val)

	if _, ok := container[pNode.Val]; !ok {
		container[pNode.Val] = make([]int, 0)

	}
	container[pNode.Val] = append(container[pNode.Val], node.Val)

	dfs6(node, node.Left, container)
	dfs6(node, node.Right, container)

}

func dfs7(pnode int, node int, container map[int][]int, dep int, k int, result *[]int) {
	if dep == k {
		*result = append(*result, node)
		return
	}

	if nodes, ok := container[node]; ok {

		for _, v := range nodes {
			if pnode == v {
				continue
			}

			dfs7(node, v, container, dep+1, k, result)
		}

	}

}
