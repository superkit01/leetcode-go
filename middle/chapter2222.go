package middle

func numberOfWays(s string) int64 {
	count0 := 0
	count1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			count0++
		} else {
			count1++
		}
	}
	prefix0 := 0
	prefix1 := 0
	var result int64 = 0
	// 当前为0时 前面1和后面1的数量乘积和
	// 当前为1时 前面0和后面0的数量乘积和
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			prefix0++
			result += int64(prefix1 * (count1 - prefix1))
		} else {
			prefix1++
			result += int64(prefix0 * (count0 - prefix0))
		}
	}
	return result

}
