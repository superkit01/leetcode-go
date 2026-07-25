package easy

func maxProduct(n int) int {
	m := []int{0, 0}
	for n > 0 {
		v := n % 10
		if v > m[0] {
			m[1] = m[0]
			m[0] = v
		} else if v > m[1] {
			m[1] = v
		}
		n /= 10
	}
	return m[0] * m[1]
}
