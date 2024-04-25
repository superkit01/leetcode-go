package middle

func countGoodNumbers(n int64) int {
	// even:=[]int{2,3,5,7}
	// odd:=[]int{0,2,4,6,8}

	if n%2 == 0 {
		return int(quickX(4, n/2) % 1000000007 * quickX(5, n/2) % 1000000007)
	} else {
		return int(quickX(4, n/2) % 1000000007 * quickX(5, n/2+1) % 1000000007)
	}

}

func quickX(x, n int64) int64 {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return quickX(x, n-1) * x % 1000000007
	}
	v := quickX(x, n/2)
	return v * v % 1000000007
}
