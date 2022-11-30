package main

//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

func countBits(n int) []int {
	result := make([]int, n+1)
	result[0] = 0
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + i&0x1
	}
	return result
}
