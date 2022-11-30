package main

func containsDuplicate(nums []int) bool {

	container := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := container[nums[i]]; !ok {
			container[nums[i]] = true
		} else {
			return true
		}
	}
	return false

}
