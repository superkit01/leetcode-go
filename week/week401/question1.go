package week401

func numberOfChild(n int, k int) int {
	k = (k + 1) % (n + n - 2)
	if k == 0 {
		return 1
	}
	if k > n {
		k = k - n
		return n - 1 - k
	} else {
		return k - 1
	}

}
