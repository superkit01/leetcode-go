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
