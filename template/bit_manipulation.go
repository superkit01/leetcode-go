package template

//将数据中的最低位的bit1清零
func removeLastBit1(a int) int {
	return a & (a - 1)
}

//是否2的幂次方
func isPowOf2(m int) bool {
	return m&(m-1) == 0
}

//lowbit ( n ) 定义为非负整数 n 在二进制表示下 “ 最低位的 1 及其后面的所有的 0 ” 的二进制构成的数值
func lowbit(a int) int {
	// pos:= math.Log2(float64(a&-a))
	return a & -a
}

//获取数据的最后一个bit0
func lastBit0(a int) int {
	// pos:= math.Log2(float64(^a & (a+1 )))
	return ^a & (a + 1)
}
