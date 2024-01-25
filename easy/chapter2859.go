package easy

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0

	for _, v := range nums {

		countBinary := 0
		binV := v

		for binV > 0 {
			if binV&0b1 == 0b1 {
				countBinary++
			}
			binV = binV >> 1
		}
		if countBinary == k {
			sum += v
		}

	}

	return sum

}
