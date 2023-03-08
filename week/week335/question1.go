package week335

func passThePillow(n int, time int) int {
	time = time % ((n - 1) * 2)

	if time <= n-1 {
		return time + 1
	} else {
		return n - (time - (n - 1))
	}
}
