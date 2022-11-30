package easy

func finalPrices(prices []int) []int {
	if len(prices) == 1 {
		return prices
	}
	result := make([]int, 0)
outer:
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				result = append(result, prices[i]-prices[j])
				continue outer
			}
		}
		result = append(result, prices[i])
	}
	result = append(result, prices[len(prices)-1])
	return result
}
