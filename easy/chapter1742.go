package easy

func countBalls(lowLimit int, highLimit int) int {
	cache := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		v := i
		for v != 0 {
			sum = sum + v%10
			v = v / 10
		}

		cache[sum]++
	}

	max := 0
	for _, v := range cache {
		if max < v {
			max = v
		}
	}
	return max
}
