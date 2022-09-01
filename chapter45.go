package main

import "math"

/**
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。

*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		steps := math.MaxInt
		for k := 0; k < i; k++ {
			if nums[k]+k >= i {
				steps = int(math.Min(float64(steps), float64(dp[k]+1)))
			}
		}
		dp[i] = steps
	}

	return dp[len(nums)-1]

}

// func main() {

// 	jump([]int{2, 3, 1, 1, 4})

// }
