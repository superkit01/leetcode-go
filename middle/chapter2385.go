package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	graph := map[int][]int{}
	dfs2385(root, &graph)

	return bfs2385(start, graph)

}

func bfs2385(start int, graph map[int][]int) int {
	//记录已经遍历过的点
	cache := make(map[int]bool, 0)
	cache[start] = true

	queue := make([]int, 0)
	queue = append(queue, start)
	ans := 0
	for len(queue) > 0 {
		ans++
		tempQueue := make([]int, 0)

		for _, c := range queue {

			for _, v := range graph[c] {
				if _, ok := cache[v]; !ok {
					tempQueue = append(tempQueue, v)
					cache[v] = true
				}
			}
		}
		queue = tempQueue
	}
	return ans - 1

}

func dfs2385(root *TreeNode, graph *map[int][]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		dfs2385(root.Left, graph)
		(*graph)[root.Val] = append((*graph)[root.Val], root.Left.Val)
		(*graph)[root.Left.Val] = append((*graph)[root.Left.Val], root.Val)
	}

	if root.Right != nil {
		dfs2385(root.Right, graph)
		(*graph)[root.Val] = append((*graph)[root.Val], root.Right.Val)
		(*graph)[root.Right.Val] = append((*graph)[root.Right.Val], root.Val)
	}

}
