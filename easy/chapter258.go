package easy

func addDigits(num int) int {
	result := num

	for result > 9 {
		result = 0
		for num > 0 {
			result += num % 10
			num = num / 10
		}
		num = result
	}
	return result

}
