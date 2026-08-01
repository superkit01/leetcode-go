package lcr

/**

1234 123

**/

func inventoryManagement(stock []int) int {

	l := 0
	r := len(stock) - 1

	for l < r {
		mid := l + (r-l)>>1
		if stock[mid] > stock[r] {
			l = mid
		} else if stock[mid] < stock[r] {
			r = mid
		} else {
			r--
		}
	}

	return stock[l]

}
