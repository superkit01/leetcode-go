package middle

func maximumSumOfHeights(maxHeights []int) int64 {
	var max int64

	for i, v := range maxHeights {
		var leftSum int64 = 0
		min := v
		for j := i; j >= 0; j-- {
			if maxHeights[j] < min {
				min = maxHeights[j]
			}
			leftSum += int64(min)
		}
		min = v
		for j := i; j < len(maxHeights); j++ {
			if maxHeights[j] < min {
				min = maxHeights[j]
			}
			leftSum += int64(min)
		}

		if max < leftSum-int64(v) {
			max = leftSum - int64(v)
		}

	}
	return max

}
