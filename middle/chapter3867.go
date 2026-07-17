package middle

import (
	"sort"
)

func gcdSum(nums []int) int64 {

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	gcdSlice := make([]int, len(nums))
	mx := nums[0]
	for i, num := range nums {
		mx = max(mx, num)
		gcdSlice[i] = gcd(num, mx)
	}

	sort.Ints(gcdSlice)

	var ans int64 = 0
	for i := 0; i < len(gcdSlice)/2; i++ {
		ans += int64(gcd(gcdSlice[i], gcdSlice[len(gcdSlice)-1-i]))
	}

	return ans

}
