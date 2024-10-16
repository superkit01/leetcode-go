package easy

import (
	"math"
	"sort"
)

func minimumAverage(nums []int) float64 {
	ans := math.MaxFloat64

	sort.Ints(nums)

	for i := 0; i < len(nums)/2; i++ {
		tmp := float64(nums[i]+nums[len(nums)-1-i]) / 2
		if ans > tmp {
			ans = tmp
		}
	}
	return ans

}
