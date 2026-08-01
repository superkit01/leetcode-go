package easy

import (
	"math"
)

func MaximumProduct(nums []int) int {

	mi := []int{math.MaxInt, math.MaxInt}
	mx := []int{math.MinInt, math.MinInt, math.MinInt}

	for _, v := range nums {

		if v > mx[0] {
			mx[2] = mx[1]
			mx[1] = mx[0]
			mx[0] = v
		} else if v > mx[1] {
			mx[2] = mx[1]
			mx[1] = v
		} else if v > mx[2] {
			mx[2] = v
		}

		if v < mi[0] {
			mi[1] = mi[0]
			mi[0] = v
		} else if v < mi[1] {
			mi[1] = v
		}

	}
	return max(mi[0]*mi[1]*mx[0], mx[0]*mx[1]*mx[2])
}
