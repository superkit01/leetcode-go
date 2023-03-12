package week336

import "sort"

func MaxScore(nums []int) int {
	positive := make([]int, 0)
	nagetive := make([]int, 0)

	for _, v := range nums {
		if v <= 0 {
			nagetive = append(nagetive, v)
		} else {
			positive = append(positive, v)
		}
	}

	if len(positive) == 0 {
		return 0
	}

	sum := 0
	for _, e := range positive {
		sum += e
	}

	sort.Ints(nagetive)

	for i := len(nagetive) - 1; i >= 0; i-- {
		if sum <= 0 {
			break
		}
		if sum+nagetive[i] <= 0 {
			break
		}

		positive = append(positive, nagetive[i])
		sum += nagetive[i]

	}

	return len(positive)

}
