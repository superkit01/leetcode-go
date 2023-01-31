package middle

import "sort"

func maxIceCream(costs []int, coins int) int {

	sort.Ints(costs)

	sum := 0
	result := 0
	for _, v := range costs {
		sum = sum + v
		if sum > coins {
			return result
		}

		result++

	}
	return result

}
