package easy

func maxProduct(nums []int) int {
	m := []int{0, 0}
	for _, v := range nums {
		if v > m[0] {
			m[1] = m[0]
			m[0] = v
		} else if v > m[1] {
			m[1] = v
		}
	}
	return (m[0] - 1) * (m[1] - 1)
}
