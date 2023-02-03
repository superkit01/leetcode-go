package middle

/**
有两位极客玩家参与了一场「二叉树着色」的游戏。游戏中，给出二叉树的根节点 root，树上总共有 n 个节点，且 n 为奇数，其中每个节点上的值从 1 到 n 各不相同。

最开始时：

「一号」玩家从 [1, n] 中取一个值 x（1 <= x <= n）；
「二号」玩家也从 [1, n] 中取一个值 y（1 <= y <= n）且 y != x。
「一号」玩家给值为 x 的节点染上红色，而「二号」玩家给值为 y 的节点染上蓝色。

之后两位玩家轮流进行操作，「一号」玩家先手。每一回合，玩家选择一个被他染过色的节点，将所选节点一个 未着色 的邻节点（即左右子节点、或父节点）进行染色（「一号」玩家染红色，「二号」玩家染蓝色）。

如果（且仅在此种情况下）当前玩家无法找到这样的节点来染色时，其回合就会被跳过。

若两个玩家都没有可以染色的节点时，游戏结束。着色节点最多的那位玩家获得胜利 ✌️。

现在，假设你是「二号」玩家，根据所给出的输入，假如存在一个 y 值可以确保你赢得这场游戏，则返回 true ；若无法获胜，就请返回 false 。

**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {

	node := findNode(root, x)

	leftCount := countNodes(node.Left)

	rightCount := countNodes(node.Right)

	other := n - 1 - leftCount - rightCount

	if other > leftCount+rightCount || leftCount > other+rightCount || rightCount > leftCount+other {
		return true
	} else {
		return false
	}

}

func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := 1
	if node.Left != nil {
		count += countNodes(node.Left)
	}
	if node.Right != nil {
		count += countNodes(node.Right)
	}
	return count
}

func findNode(node *TreeNode, x int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == x {
		return node
	}
	target := findNode(node.Left, x)
	if target != nil {
		return target
	} else {
		return findNode(node.Right, x)
	}
}
