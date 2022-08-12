package main

func groupThePeople(groupSizes []int) [][]int {

	cache := make(map[int][]int, 0)

	for k, v := range groupSizes {
		if _, ok := cache[v]; !ok {
			cache[v] = make([]int, 0)
		}
		cache[v] = append(cache[v], k)
	}
	result := make([][]int, 0)
	for k, v := range cache {
		if k == len(v) {
			result = append(result, v)
		} else {
			count := len(v) / k

			for i := 0; i < count; i++ {
				result = append(result, v[i*k:(i+1)*k])
			}
		}
	}
	return result

}
