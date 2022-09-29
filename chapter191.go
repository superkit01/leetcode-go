package main


func hammingWeight(num uint32) int {
	result := 0
	for num > 0 {

		if num&0b1 == 1 {
			result += 1
		}
		num = num >> 1
	}

	return result

}
