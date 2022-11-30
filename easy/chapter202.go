package main

import "math"

func isHappy(n int) bool {
	cache := make(map[int]bool)
	temp := n
	for {
		if temp == 1 {
			return true
		}

		if _, ok := cache[temp]; ok {
			return false
		}

		cache[temp] = true

		newTemp := 0

		for temp != 0 {
			newTemp += int(math.Pow(float64(temp%10), 2))
			temp = temp / 10
		}

		temp = newTemp

	}

}
