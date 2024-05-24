package middle

func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	lNode := lowestCommonAncestor(root.Left, p, q)
	rNode := lowestCommonAncestor(root.Right, p, q)
	if lNode != nil && rNode != nil {
		return root
	}
	if rNode == nil {
		return lNode
	}
	return rNode

}
