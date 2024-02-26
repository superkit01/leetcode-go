package middle

func NumOfSubarrays(arr []int, k int, threshold int) int {
	sum := 0
	result := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	if sum >= threshold*k {
		result++
	}

	for i := k; i < len(arr); i++ {
		sum = sum + arr[i] - arr[i-k]
		if sum >= threshold*k {
			result++
		}

	}
	return result

}
