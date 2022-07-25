package main

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	cache := make(map[int]bool)

	for _, v := range nums1 {
		if _, ok := cache[v]; !ok {
			cache[v] = true
		}
	}

	for _, v := range nums2 {
		if _, ok := cache[v]; ok {
			result = append(result, v)
			delete(cache, v)
		}
	}

	return result

}
