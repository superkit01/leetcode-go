package middle

import "strings"

// 10 --> 01  后置的0前移
// 00 --> 10  10 -->01  ===   00 --> 01   最终仅剩余1个0

// 01111110111111111111

func MaximumBinaryString(binary string) string {

	count0 := 0

	for _, v := range binary {
		if v == '0' {
			count0++
		}
	}

	if count0 <= 1 {
		return binary
	}

	ans := ""
	start := -1

	for i, v := range binary {
		if v == '0' {
			start = i
			break
		}
	}

	ans += strings.Repeat("1", start+(count0-1))
	ans += "0"
	ans += strings.Repeat("1", len(binary)-(start+count0))

	return ans

}
