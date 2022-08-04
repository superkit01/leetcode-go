package main

import (
	"sort"
)

func minSubsequence(nums []int) []int {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	sort.Ints(nums)

	temp := 0

	for i := len(nums) - 1; i >= 0; i-- {
		temp += nums[i]

		if temp >= sum/2+1 {
			result := nums[i:]

			sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

			return result
		}

	}
	return []int{0}

}

func main() {
	minSubsequence([]int{1})

}
