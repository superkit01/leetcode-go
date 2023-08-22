package middle

func maxDistToClosest(seats []int) int {
	current := -1

	prefixSum := make([]int, 0)

	for i, v := range seats {
		if v == 1 {
			current = i
			prefixSum = append(prefixSum, 0)
		} else {
			prefixSum = append(prefixSum, i-current)
		}
	}

	max := 0
	for _, v := range prefixSum {
		if max < v {
			max = v
		}
	}

	left := 0
	for _, v := range prefixSum {
		if v == 0 {
			break
		} else {
			left = v
		}
	}

	right := 0
	for i := len(prefixSum) - 1; i > 0; i-- {
		if prefixSum[i] == 0 {
			break
		} else {
			right = prefixSum[i]
			break
		}
	}

	if right > left {
		left = right
	}

	if max%2 == 0 {
		max = max / 2
	} else {
		max = (max + 1) / 2
	}

	if left > max {
		return left
	} else {
		return max
	}

}
