package lcr

func mechanicalAccumulator(target int) int {
	ans := 0
	var acc func(n int) bool
	acc = func(n int) bool {
		ans += n
		return n > 0 && acc(n-1)
	}
	acc(target)
	return ans

}
