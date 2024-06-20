package easy

func CountBeautifulPairs(nums []int) int {
	cnt := map[int]int{}

	ans := 0
	for _, n := range nums {
		last := n % 10
		for k, v := range cnt {
			if gcd(k, last) == 1 {
				ans += v
			}
		}

		for n >= 10 {
			n = n / 10
		}
		cnt[n]++

	}
	return ans

}

func gcd(x, y int) int {
	if x%y == 0 {
		return y
	}
	return gcd(y, x%y)

}
