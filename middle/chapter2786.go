package middle

import "math"

func maxScore(nums []int, x int) int64 {
	dp := make([]int64, len(nums))
	even := -1
	odd := -1
	dp[0] = int64(nums[0])
	if nums[0]%2 == 1 {
		odd = 0
	} else {
		even = 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]%2 == nums[i-1]%2 {
			dp[i] = dp[i-1] + int64(nums[i])
			if nums[i]%2 == 1 {
				odd = odd + 1
			} else {
				even = even + 1
			}

		} else {
			if nums[i]%2 == 1 {
				if nums[odd]%2 == nums[i]%2 {
					dp[i] = max(dp[i-1]-int64(x), dp[odd]) + int64(nums[i])
				} else {
					dp[i] = dp[i-1] + int64(nums[i]) - int64(x)
				}
				odd = i
			} else {
				if nums[even]%2 == nums[i]%2 {
					dp[i] = max(dp[i-1]-int64(x), dp[even]) + int64(nums[i])
				} else {
					dp[i] = dp[i-1] + int64(nums[i]) - int64(x)
				}
				even = i
			}

		}

	}

	return dp[len(dp)]
}

func max[T int|int64](i, j T) T {
	if i > j {
		return i
	}
	return j
}

func maxScoreII(nums []int, x int) int64 {
	dp := make([][2]int64, len(nums))
	//[[偶][奇]]
	dp[0] = [2]int64{math.MinInt32, math.MinInt32}
	dp[0][nums[0]%2] = int64(nums[0])
	for i := 1; i < len(nums); i++ {
		dp[i] = [2]int64{dp[i-1][0], dp[i-1][1]}
		if nums[i]%2 == 0 {
			dp[i][0] = max(dp[i-1][0]+int64(nums[i]), dp[i-1][1]+int64(nums[i])-int64(x))
		} else {
			dp[i][1] = max(dp[i-1][1]+int64(nums[i]), dp[i-1][0]+int64(nums[i])-int64(x))
		}
	}
	return max(dp[len(dp)-1][0], dp[len(dp)-1][1])
}
