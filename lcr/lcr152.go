package lcr

// |左节点|右节点|根节点|
func verifyTreeOrder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	root := postorder[len(postorder)-1]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; i < len(postorder); i++ {
		if postorder[i] >= root {
			left = postorder[0:i]
			right = postorder[i : len(postorder)-1]
			break
		}
	}

	for m := 0; m < len(right); m++ {
		if right[m] <= root {
			return false
		}
	}

	return verifyTreeOrder(left) && verifyTreeOrder(right)

}
