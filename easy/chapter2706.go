package easy

import "sort"

func buyChoco(prices []int, money int) int {

	sort.Ints(prices)

	if prices[0]+prices[1] > money {
		return money
	} else {
		return money - prices[0] - prices[1]
	}

}
