package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {

	result := make([]int, 0)
	if root == nil {
		return result
	}

	cache := make([]*TreeNode, 0)
	cache = append(cache, root)
	result = append(result, root.Val)

	for len(cache) != 0 {
		temp := make([]*TreeNode, 0)
		for i := 0; i < len(cache); i++ {
			if cache[i].Left != nil {
				temp = append(temp, cache[i].Left)
			}
			if cache[i].Right != nil {
				temp = append(temp, cache[i].Right)
			}
		}
		if len(temp) != 0 {
			result = append(result, temp[len(temp)-1].Val)
		}
		cache = temp
	}
	return result
}
