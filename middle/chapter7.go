package middle

func reverse(x int) int {
	var res int64 = 0
	for x != 0 {
		res = res*int64(10) + int64(x%10)
		x = x / 10
	}
	if int64(int32(res)) == res {
		return int(res)
	} else {
		return 0
	}
}
