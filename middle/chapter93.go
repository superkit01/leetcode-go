package middle

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)

	if len(s) < 4 || len(s) > 12 {
		return result
	}

	digui3(s, 3, "", &result)
	return result
}

func digui3(s string, n int, temp string, result *[]string) {

	if n == 0 {
		if len(s) > 3 {
			return
		}
		value, _ := strconv.Atoi(s)
		if value > 255 {
			return
		}
		if strings.HasPrefix(s, "0") && s != "0" {
			return
		}
		temp = temp + s
		*result = append(*result, temp)
	} else {
		for i := 1; i <= 3; i++ {

			if len(s) <= i {
				return
			}
			currS := s[0:i]

			value, _ := strconv.Atoi(currS)
			if value > 255 {
				return
			}
			if strings.HasPrefix(currS, "0") && currS != "0" {
				return
			}

			temps := s[i:]
			digui3(temps, n-1, temp+strconv.Itoa(value)+".", result)
		}
	}

}

// func main() {

// 	restoreIpAddresses("010010")
// }
