package easy

import "math"

func average(salary []int) float64 {
	if len(salary) <= 2 {
		return 0
	}
	max := 0
	min := math.MaxInt32
	sum := 0

	for _, v := range salary {
		sum += v
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return float64(sum-min-max) / float64(len(salary)-2)

}
