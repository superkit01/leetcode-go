package main

func numberOfPairs(nums []int) []int {
	result := [2]int{0, len(nums)}

	container := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if v, ok := container[nums[i]]; !ok || !v {
			container[nums[i]] = true
		} else {
			container[nums[i]] = false
			result[0] += 1
			result[1] -= 2
		}

	}
	return result[:]

}
