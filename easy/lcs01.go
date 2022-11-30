package main

func leastMinutes(n int) int {
	if n == 1 {
		return 1
	}

	result := 1
	count := 0
	for result < n {
		result = 2 * result
		count += 1
	}
	return count + 1

}
