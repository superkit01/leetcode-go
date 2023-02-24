package easy

func minimumOperations(nums []int) int {
	container := make(map[int]bool, 0)
	for _, v := range nums {
		if v == 0 {
			continue
		} else {
			container[v] = true
		}
	}

	return len(container)
}
