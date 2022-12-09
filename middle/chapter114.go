package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	cache := make([]*TreeNode, 0)

	cache = preTraversal(root, cache)

	if len(cache) > 1 {
		for i := 0; i < len(cache)-1; i++ {
			cache[i].Left = nil
			cache[i].Right = cache[i+1]
		}
		cache[len(cache)-1].Left = nil
	}

	root = cache[0]

}

func preTraversal(node *TreeNode, cache []*TreeNode) []*TreeNode {
	if node == nil {
		return cache
	}
	cache = append(cache, node)
	cache = preTraversal(node.Left, cache)
	cache = preTraversal(node.Right, cache)
	return cache
}
