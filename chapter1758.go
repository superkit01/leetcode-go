package main

import "math"

func minOperations1(s string) int {
	result := 0
	flag := '0'
	for _, v := range s {
		if v != flag {
			result++
		}

		if flag == '1' {
			flag = '0'
		} else {
			flag = '1'
		}
	}

	return int(math.Min(float64(result), float64(len(s)-result)))

}
