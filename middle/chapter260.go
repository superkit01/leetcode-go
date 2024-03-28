package middle

func singleNumber(nums []int) []int {
	z := 0
	for _, v := range nums {
		z = v ^ z
	}

	//  1  ...000001
    // -1  ...111111
    //最后一位二进制1
	pos1 := z & -z

	x0 := 0
	x1 := 0
	for _, v := range nums {
		if (pos1 & v) == pos1 {
			x0 = x0 ^ v
		} else {
			x1 = x1 ^ v
		}
	}

	return []int{x0, x1}
}
