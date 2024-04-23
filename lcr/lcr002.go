package lcr

import "fmt"

func AddBinary(a string, b string) string {
	i := 0
	ans := ""
	c := 0
	for i < len(a) && i < len(b) && c == 0 {

		aValue := 0
		if i < len(a) {
			aValue = int(a[len(a)-1-i] - '0')
		}

		bValue := 0
		if i < len(b) {
			bValue = int(b[len(b)-1-i] - '0')
		}
		divide := (aValue + bValue + c) % 2
		c = (aValue + bValue + c) / 2

		ans = fmt.Sprint(divide) + ans
		i++
	}
	return ans

}
