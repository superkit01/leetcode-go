package easy

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {

	if original == target {
		return cloned
	}
	var ans *TreeNode
	if original.Left != nil {
		ans = getTargetCopy(original.Left, cloned.Left, target)

	}
	if ans == nil {
		if original.Right != nil {
			ans = getTargetCopy(original.Right, cloned.Right, target)
		}
	}

	return ans

}
