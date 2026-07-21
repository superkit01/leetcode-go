package lcr

import (
	"bytes"
	"sort"
	"strconv"
)

/**

[0, 3, 30, 34, 5, 9]

**/

func crackPassword(password []int) string {
	passStr := make([]string, len(password))
	for i := range passStr {
		passStr[i] = strconv.Itoa(password[i])
	}

	sort.Slice(passStr, func(i, j int) bool {
		if passStr[i]+passStr[j] > passStr[j]+passStr[i] {
			return false
		} else {
			return true
		}
	})

	var buf bytes.Buffer
	for _, v := range passStr {
		buf.WriteString(v)
	}
	return buf.String()
}
