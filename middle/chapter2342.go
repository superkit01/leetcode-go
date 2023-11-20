package middle

import "math"

func maximumSum(nums []int) int {
	cache := make(map[int][]int, 0)

	for _, v := range nums {
		sum := 0
		temp := v
		for temp > 0 {
			sum += temp % 10
			temp = temp / 10
		}

		if _, ok := cache[sum]; !ok {
			cache[sum] = []int{0, 0}
		}

		if cache[sum][1] <= v {
			cache[sum][0] = cache[sum][1]
			cache[sum][1] = v

		} else if cache[sum][0] <= v {
			cache[sum][0] = v
		}
	}

	max := -1
	for _, v := range cache {
		if v[0] == 0 {
			continue
		}
		max = int(math.Max(float64(v[0])+float64(v[1]), float64(max)))
	}

	return max
}
