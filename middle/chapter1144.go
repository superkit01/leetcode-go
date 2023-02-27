package middle

import "math"

func movesToMakeZigzag(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	caseEven := 0

	for i := 1; i < len(nums); i += 2 {
		if i >= len(nums) {
			break
		}
		if i == len(nums)-1 {
			if nums[i] >= nums[i-1] {
				caseEven += nums[i] - nums[i-1] + 1
			}
		} else {
			if nums[i] >= nums[i-1] || nums[i] >= nums[i+1] {
				caseEven += nums[i] - int(math.Min(float64(nums[i-1]), float64(nums[i+1]))) + 1
			}
		}

	}

	caseOdd := 0

	for i := 0; i < len(nums); i += 2 {
		if i >= len(nums) {
			break
		}
		if i == len(nums)-1 {
			if nums[i] >= nums[i-1] {
				caseOdd += nums[i] - nums[i-1] + 1
			}
		} else if i == 0 {
			if nums[i] >= nums[i+1] {
				caseOdd += nums[i] - nums[i+1] + 1
			}

		} else {
			if nums[i] >= nums[i-1] || nums[i] >= nums[i+1] {
				caseOdd += nums[i] - int(math.Min(float64(nums[i-1]), float64(nums[i+1]))) + 1
			}
		}

	}

	return int(math.Min(float64(caseEven), float64(caseOdd)))

}
