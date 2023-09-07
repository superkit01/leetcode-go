package middle

import (
	"math"
	"sort"
)

func RepairCars(ranks []int, cars int) int64 {
	sort.Ints(ranks)

	left := 1
	right := ranks[len(ranks)-1] * (cars/len(ranks) + 1) * (cars/len(ranks) + 1)
	for left < right {
		mid := (left + right) / 2
		sum := 0
		for _, v := range ranks {
			count := int(math.Sqrt(float64(mid / v)))
			sum += count
		}
		if sum >= cars {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return int64(left)
}
