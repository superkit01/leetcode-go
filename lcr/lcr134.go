package lcr

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n%2 == 1 {
		return myPow(x, n-1) * x
	} else {
		temp := myPow(x, n/2)
		return temp * temp
	}

}
