package main

import "math"

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	result := 0
	for _, v := range fruits {
		count := math.Ceil(float64(v[1]) / float64(limit))
		perTime := time[v[0]]
		result += int(count) * perTime
	}
	return result

}
