package hard

func longestAwesome(s string) int {
	ans := 0

	cache := [1 << 10]int{}
	for i := range cache {
		cache[i] = -1
	}

	bitPrefix := 0

outer:
	for i, v := range s {
		bitPrefix ^= 1 << (v - '0')
		//3242415
		// 100  110  1110  1100  0100  0101  10101
		// 0	1	 2	   3     4     5     6
		if cache[bitPrefix] == -1 {
			cache[bitPrefix] = i
		}

		//  bit & bit-1 == 0 ==> 二进制表示只有一个1
		if bitPrefix == 0 || bitPrefix&(bitPrefix-1) == 0 {
			ans = max(ans, i+1)
			continue outer
		}

		// ans = max(ans, i-cache[bitPrefix])
		for j := 0; j < 10; j++ {
			temp := bitPrefix ^ (1 << j)
			if cache[temp] != -1 {
				ans = max(ans, i-cache[temp])
			}
		}
	}

	return ans

}
