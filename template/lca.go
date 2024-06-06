package template

import (
	"math"
)

//lowest common ancesters 最近公共祖先

// dfs 深度优先 递归
func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lca(root.Left, p, q)
	right := lca(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left

}

// lca一定在数组的p、q节点对应位置的中间 且深度最低
func lcaII(root, p, q *TreeNode) *TreeNode {
	inOrder := make([]*DeepNode, 0)

	//带深度的中序遍历 |左|跟|右|
	dfs(root, &inOrder, 1)

	pIndex := -1
	qIndex := -1
	for i := 0; i < len(inOrder); i++ {
		if inOrder[i].node == p {
			pIndex = i
		}
		if inOrder[i].node == q {
			qIndex = i
		}
	}

	if pIndex > qIndex {
		pIndex, qIndex = qIndex, pIndex
	}

	maxDeep := math.MaxInt32
	var ans *TreeNode
	for i := pIndex; i <= qIndex; i++ {
		if maxDeep > inOrder[i].deep {
			maxDeep = inOrder[i].deep
			ans = inOrder[i].node
		}
	}

	return ans

}

func dfs(root *TreeNode, inOrder *[]*DeepNode, deep int) {
	if root == nil {
		return
	}
	dfs(root.Left, inOrder, deep+1)
	*inOrder = append(*inOrder, &DeepNode{
		node: root,
		deep: deep,
	})
	dfs(root.Right, inOrder, deep+1)

}

type DeepNode struct {
	node *TreeNode
	deep int
}
