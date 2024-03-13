package easy

import "strings"

func maximumOddBinaryNumber(s string) string {
	count1 := 0
	for _, v := range s {
		if v == '1' {
			count1++
		}
	}
	ans := ""

	if count1 > 0 {
		ans = ans + strings.Repeat("0", len(s)-count1) + "1"
		count1--
	}
	if count1 > 0 {
		ans = strings.Repeat("1", count1) + ans
	}

	return ans
}
