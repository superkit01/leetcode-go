package easy

func countDigits(num int) int {
	current := num
	result := 0

	for current != 0 {
		if num%(current%10) == 0 {
			result++
		}
		current = current / 10

	}

	return result

}
