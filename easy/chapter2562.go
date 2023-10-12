package easy

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	left := 0
	right := len(nums) - 1

	var result int64 = 0
	for left < right {
		v, _ := strconv.ParseInt(strconv.Itoa(nums[left])+strconv.Itoa(nums[right]), 10, 64)
		result += v
		left++
		right--
	}

	if left == right {
		result += int64(nums[left])
	}

	return int64(result)

}
