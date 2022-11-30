package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	container := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := container[nums[i]]; !ok {
			container[nums[i]] = i
		} else {
			if math.Abs(float64(container[nums[i]]-i)) <= float64(k) {
				return true
			} else {
				container[nums[i]] = i
			}
		}
	}
	return false

}
