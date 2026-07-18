package middle

func rangeBitwiseAnd(left int, right int) int {
	// 二进制位的最长公共前缀

	// 		 z   1110
	//		z-1	 1101
	// 	 z & z-1 1100   去掉最低位的1
	//       z-1 1011
	//   z & z-1 1000
	//  去掉

	for left < right {
		right = right & (right - 1)
	}
	return right

}
