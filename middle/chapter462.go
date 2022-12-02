package middle

import "sort"

func MinMoves2(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	sort.Ints(nums)

	if len(nums)%2 == 0 {
		sum := 0
		for i := 0; i < len(nums)/2; i++ {
			sum += nums[len(nums)-1-i] - nums[i]
		}
		return sum
	} else {
		sum := 0
		for i := 0; i < len(nums)/2; i++ {
			sum += nums[len(nums)-1-i] - nums[i]
		}
		return sum
	}

}
