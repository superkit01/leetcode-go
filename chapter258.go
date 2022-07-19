package main

func addDigits(num int) int {
	result := num
	
1	for result > 9 {
		result := 0
		for num > 0 {
			result += num % 10
			num = num / 10
		}
		num = result
	}
	return result

}
