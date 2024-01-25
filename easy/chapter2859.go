package easy

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		countBinary := 0
		binV := i

		for binV > 0 {
			if binV&0b1 == 0b1 {
				countBinary++
			}
			binV = binV >> 1
		}
		if countBinary == k {
			sum += nums[i]
		}
	}

	return sum

}
