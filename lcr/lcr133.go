package lcr

func hammingWeight(num uint32) int {
	ans := 0
	for i := 0; i < 32; i++ {
		if num&(1<<i) == (1 << i) {
			ans++
		}
	}
	return ans

}

func hammingWeightII(num uint32) int {
	ans := 0
	for num != 0 {
		num = num & (num - 1) // bit & bit-1 去除二进制位的最后一个1
		ans++
	}
	return ans
}
