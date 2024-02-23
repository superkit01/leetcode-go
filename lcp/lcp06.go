package lcp

func minCount(coins []int) int {
	sum := 0

	for _, v := range coins {
		if v%2 == 1 {
			sum += 1
		}
		sum += v >> 1
	}
	return sum
}
