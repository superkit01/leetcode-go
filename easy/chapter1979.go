package easy

func findGCD(nums []int) int {
	mi, mx := nums[0], nums[0]

	for _, v := range nums {
		mi = min(mi, v)
		mx = max(mx, v)
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	return gcd(mi, mx)
}
