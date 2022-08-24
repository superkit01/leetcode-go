package main

func canBeEqual(target []int, arr []int) bool {
	cache := make(map[int]int, 0)
	for i := range target {
		cache[target[i]]++
		cache[arr[i]]--
	}

	for _, v := range cache {
		if v != 0 {
			return false
		}
	}
	return true

}
