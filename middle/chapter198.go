package middle

import "math"

func rob(nums []int) int {
	take := make([]int, len(nums))
	noTake := make([]int, len(nums))

	take[0] = nums[0]
	noTake[0] = 0

	if len(nums) == 1 {
		return take[0]
	}

	for i := 1; i < len(nums); i++ {
		take[i] = noTake[i-1] + nums[i]
		noTake[i] = int(math.Max(float64(take[i-1]), float64(noTake[i-1])))
	}

	if take[len(nums)-1] > noTake[len(nums)-1] {
		return take[len(nums)-1]
	} else {
		return noTake[len(nums)-1]
	}

}
