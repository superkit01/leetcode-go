package week366

import "sort"

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(processorTime)

	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))

	result := make([]int, 0)

	for i := 0; i < len(processorTime); i++ {
		max := tasks[0+4*i]
		if tasks[1+4*i] > max {
			max = tasks[1+4*i]
		}
		if tasks[2+4*i] > max {
			max = tasks[2+4*i]
		}
		if tasks[3+4*i] > max {
			max = tasks[3+4*i]
		}
		result = append(result, max+processorTime[i])
	}

	maxValue := 0

	for i := 0; i < len(result); i++ {
		if maxValue < result[i] {
			maxValue = result[i]
		}
	}
	return maxValue

}
