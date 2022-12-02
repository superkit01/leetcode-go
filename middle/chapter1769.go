package middle

func minOperations(boxes string) []int {
	if len(boxes) == 1 {
		return []int{0}
	}

	queueRight := make([]int, 0)
	queueLeft := make([]int, 0)
	result := make([]int, 0)
	for i, v := range boxes {
		if v == '1' {
			if i == 0 {
				queueLeft = append(queueLeft, i)
			} else {
				queueRight = append(queueRight, i)
			}
		}
	}

	sum := 0
	for _, v := range queueRight {
		sum += v
	}

	result = append(result, sum)

	for i := 1; i < len(boxes); i++ {
		last := result[len(result)-1]
		result = append(result, last+len(queueLeft)-len(queueRight))

		if len(queueRight) > 0 && queueRight[0] == i {
			queueLeft = append(queueLeft, queueRight[0])
			queueRight = queueRight[1:]
		}

	}
	return result

}
