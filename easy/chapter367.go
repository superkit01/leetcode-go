package main

func isPerfectSquare(num int) bool {
	left := 1
	right := num

	for left <= right {
		mid := (right + left) / 2
		square := mid * mid
		if square > num {
			right = mid - 1
		} else if square < num {
			left = mid + 1
		} else {
			return true
		}

	}
	return false
}
