package middle

import "math"

func maxProduct(words []string) int {
	cache := make([]int, 0)

	for _, v := range words {
		binaryValue := 0
		for _, c := range v {
			binaryValue |= 1 << (c - 'a')
		}
		cache = append(cache, binaryValue)
	}
	max := 0
	for i := 0; i < len(cache)-1; i++ {
		for j := i; j < len(cache); j++ {
			if cache[i]&cache[j] == 0 {
				max = int(math.Max(float64(len(words[i])*len(words[j])), float64(max)))
			}

		}
	}
	return max

}
