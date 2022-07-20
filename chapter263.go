package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	remander := n
	for remander != 1 {
		if remander%2 == 0 {
			remander = remander / 2
		} else if remander%3 == 0 {
			remander = remander / 3
		} else if remander%5 == 0 {
			remander = remander / 5
		} else {
			return false
		}
	}
	return true

}
