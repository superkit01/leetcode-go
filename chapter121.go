package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	minPrice := prices[0]

	maxResult := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if maxResult < prices[i]-minPrice {
			maxResult = prices[i] - minPrice
		}

	}
	return maxResult

}
