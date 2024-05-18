package easy

import "math"

func maxDivScore(nums []int, divisors []int) int {
	cnt := make([]int, len(divisors))
	for i, v := range divisors {
		for _, k := range nums {
			if k%v == 0 {
				cnt[i]++
			}
		}
	}
	maxIndex := -1
	maxValue := math.MinInt32
	for i, v := range cnt {
		if maxValue < v {
			maxValue = v
			maxIndex = i
		} else if maxValue == v {
			if divisors[maxIndex] > divisors[i] {
				maxIndex = i
			}
		}
	}

	return divisors[maxIndex]
}
