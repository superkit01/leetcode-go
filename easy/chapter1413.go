package easy

func minStartValue(nums []int) int {
	temp := 0

	prefixSum := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		prefixSum[i] = temp
	}

	minValue := prefixSum[0]

	for i := 0; i < len(prefixSum); i++ {
		if minValue > prefixSum[i] {
			minValue = prefixSum[i]
		}
	}
	if minValue >= 1 {
		return 1
	} else {
		return 1 - minValue
	}
}
