package lcr

import "math"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	inOrder := make([]*DeepNode, 0)

	//带深度的中序遍历 |左|跟|右|
	dfsLCA(root, &inOrder, 1)

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

func dfsLCA(root *TreeNode, inOrder *[]*DeepNode, deep int) {
	if root == nil {
		return
	}
	dfsLCA(root.Left, inOrder, deep+1)
	*inOrder = append(*inOrder, &DeepNode{
		node: root,
		deep: deep,
	})
	dfsLCA(root.Right, inOrder, deep+1)

}

type DeepNode struct {
	node *TreeNode
	deep int
}
