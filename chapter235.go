package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}
	if root.Val <= p.Val && root.Val >= q.Val {
		return root
	}
	return nil

}
