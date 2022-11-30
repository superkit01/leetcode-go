package main

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	cache := make(map[int]int)

	for _, v := range nums1 {
		if _, ok := cache[v]; !ok {
			cache[v] = 0
		}
		cache[v]++
	}

	for _, v := range nums2 {
		if c, ok := cache[v]; ok {
			if c > 0 {
				result = append(result, v)
				cache[v]--
			}
		}
	}

	return result
}
