package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	root *TreeNode

	cache map[int]bool
}

func Constructor3(root *TreeNode) FindElements {

	cache := make(map[int]bool, 0)
	tree := FindElements{
		root:  root,
		cache: cache,
	}

	root.Val = 0
	cache[0] = true

	dfs(root, cache)

	return tree

}

func dfs(root *TreeNode, cache map[int]bool) {
	if root.Left != nil {
		root.Left.Val = (root.Val)*2 + 1
		cache[root.Left.Val] = true
		dfs(root.Left, cache)
	}

	if root.Right != nil {
		root.Right.Val = (root.Val)*2 + 2
		cache[root.Right.Val] = true
		dfs(root.Right, cache)
	}
}

func (this *FindElements) Find(target int) bool {
	if _, ok := this.cache[target]; ok {
		return true
	} else {
		return false
	}
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
