package lcr

import "math"

func bestTiming(prices []int) int {
	minonStack := make([]int, 0)

	ans := 0

outer:
	for i := len(prices) - 1; i >= 0; i-- {

		for {
			if len(minonStack) == 0 {
				minonStack = append(minonStack, prices[i])
				continue outer
			}

			if prices[i] < minonStack[len(minonStack)-1] {
				minonStack = append(minonStack, prices[i])
				ans = int(math.Max(float64(ans), float64(minonStack[0]-minonStack[len(minonStack)-1])))
				continue outer
			}

			if prices[i] >= minonStack[len(minonStack)-1] {
				minonStack = minonStack[0 : len(minonStack)-1]
			}
		}

	}

	return ans

}

func bestTimingII(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := prices[0]
	ans := 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - min //股票在第i天卖出的利润
		ans = int(math.Max(float64(ans), float64(profit)))

		min = int(math.Min(float64(min), float64(prices[i])))
	}
	return ans

}
