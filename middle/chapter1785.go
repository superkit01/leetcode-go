package middle

import "math"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum == goal {
		return 0
	}

	dis := int(math.Abs(float64(goal - sum)))

	if dis%limit == 0 {
		return dis / limit
	} else {
		return dis/limit + 1
	}
}
