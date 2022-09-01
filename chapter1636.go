package main

import "sort"

func frequencySort(nums []int) []int {
	cache := make(map[int]int, 0)

	for _, v := range nums {
		if _, ok := cache[v]; !ok {
			cache[v] = 0
		}
		cache[v]++
	}

	cache2 := make(map[int][]int, 0)
	keys := make([]int, 0)
	for k, v := range cache {
		if _, ok := cache2[v]; !ok {
			cache2[v] = make([]int, 0)
			keys = append(keys, v)
		}
		cache2[v] = append(cache2[v], k)
	}

	sort.Ints(keys)

	result := make([]int, 0)
	for _, v := range keys {
		temp := cache2[v]
		sort.Sort(sort.Reverse(sort.IntSlice(temp)))
		for _, e := range temp {
			for i := 0; i < v; i++ {
				result = append(result, e)
			}
		}
	}

	return result
}

// func main() {
// frequencySort([]int{1, 1, 2, 2, 2, 3})
// }
