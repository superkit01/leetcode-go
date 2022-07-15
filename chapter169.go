package main

func majorityElement(nums []int) int {
	counter := make(map[int]int, 0)

	n := len(nums)

	for i := 0; i < n; i++ {
		if _, ok := counter[nums[i]]; !ok {
			counter[nums[i]] = 0
		}
		counter[nums[i]]++
		if counter[nums[i]] > n/2 {
			return nums[i]
		}
	}
	return 0
}
