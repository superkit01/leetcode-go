package main

func arraySign(nums []int) int {
	result := 1
	for _,i := range nums {
		if i == 0 {
			return 0
		} else if i < 0 {
			result *= -1
		}
	}
	return result
}
