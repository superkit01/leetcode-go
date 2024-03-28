package lcs

import "sort"

func halfQuestions(questions []int) int {
	n := len(questions) / 2

	countMap := make(map[int]int)
	for i := 0; i < len(questions); i++ {
		if _, ok := countMap[questions[i]]; !ok {
			countMap[questions[i]] = 0
		}
		countMap[questions[i]]++
	}

	slice := make([]int, 0)

	for _, v := range countMap {
		slice = append(slice, v)
	}

	sort.Ints(slice)

	sum := 0
	count := 0
	for i := len(slice) - 1; i >= 0; i-- {
		if sum < n {
			sum += slice[i]
			count += 1
		} else {
			return count
		}

	}

	return count

}
