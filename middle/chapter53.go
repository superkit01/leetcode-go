package middle

import "math"

func maxSubArray(nums []int) int {

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max

}

func maxSubArray1(nums []int) int {

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		max = int(math.Max(float64(nums[i]), float64(max)))

	}

	return max
}
