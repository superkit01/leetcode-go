package main

import "sort"

func arrayRankTransform(arr []int) []int {

	temp := make([]int, 0)
	temp = append(temp, arr...)
	sort.Ints(temp)

	cache := make(map[int]int, 0)
	index := 1
	for _, v := range temp {
		if _, ok := cache[v]; !ok {
			cache[v] = index
			index++
		}
	}

	result := make([]int, 0)
	for _, v := range arr {
		result = append(result, cache[v])
	}
	return result

}

// func main() {
// 	arrayRankTransform([]int{40, 10, 20, 30})
// }
