package middle

import "bytes"

func addSpaces(s string, spaces []int) string {
	var buffer bytes.Buffer
	j := 0
	for i, v := range s {
		if len(spaces) > j && i == spaces[j] {
			buffer.WriteRune(' ')
			j++
		}
		buffer.WriteRune(v)
	}
	return buffer.String()

}
