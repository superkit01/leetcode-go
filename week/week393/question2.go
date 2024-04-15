package week393

import "math"

func maximumPrimeDifference(nums []int) int {
	start := -1

	for i := 0; i < len(nums); i++ {
		if isPrime(nums[i]) {
			start = i
			break
		}
	}

	end := -1

	for i := len(nums) - 1; i >= 0; i-- {
		if isPrime(nums[i]) {
			end = i
			break
		}
	}

	return end - start

}

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	if x == 2 {
		return true
	}

	right := int(math.Sqrt(float64(x)))

	for i := 2; i <= right; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
