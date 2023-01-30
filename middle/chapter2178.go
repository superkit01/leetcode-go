package middle

func maximumEvenSplit(finalSum int64) []int64 {
	result := make([]int64, 0)
	if finalSum%2 != 0 {
		return result
	}

	var temp int64 = 0
	for {
		temp = temp + 2

		if finalSum <= 2*temp {
			break
		}
		result = append(result, int64(temp))

		finalSum = finalSum - int64(temp)
	}

	result = append(result, finalSum)
	return result
}
