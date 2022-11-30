package hard

func nthMagicalNumber(n int, a int, b int) int {
	if a > b {
		a, b = b, a
	}
	v := gcd1(b, a)
	gbs := a * b / v
	right := a * n
	left := a
	for left <= right {

		mid := (left + right) / 2

		count := mid/a + mid/b - mid/gbs

		if count >= n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return (right + 1) % (1e9 + 7)

}

func gcd1(a int, b int) int {
	mod := a % b
	if mod == 0 {
		return b
	} else {
		return gcd1(b, mod)
	}
}
