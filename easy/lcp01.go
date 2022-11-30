package main

func game(guess []int, answer []int) int {
	result := 0
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			result++
		}

	}
	return result

}
