package main

func mySqrt(x int) int {
	left := 1
	right := x

	for left <= right {
		mid := (left + right) / 2

		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}

	return right
}

//func main() {
//	fmt.Printf("%v  \n", mySqrt(16))
//}
