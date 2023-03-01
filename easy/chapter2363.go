package easy

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	container := make(map[int]int, 0)
	for _, v := range items1 {
		container[v[0]] = container[v[0]] + v[1]
	}
	for _, v := range items2 {
		container[v[0]] = container[v[0]] + v[1]
	}
	result := make([][]int, 0)
	for k, v := range container {
		result = append(result, []int{k, v})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result

}
