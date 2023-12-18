package easy

import "math"

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return int(math.Min(float64(cost[0]), float64(cost[1])))
	}

	for i := 2; i < len(cost); i++ {
		cost[i] = int(math.Min(float64(cost[i-1]), float64(cost[i-2]))) + cost[i]
	}

	return int(math.Min(float64(cost[len(cost)-1]), float64(cost[len(cost)-2])))

}
