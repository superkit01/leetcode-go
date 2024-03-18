package middle

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	sum := 0
	for i := 1; i < len(prices); i++ {
		sum += int(math.Max(float64(prices[i]-prices[i-1]), 0))
	}
	return sum

}
