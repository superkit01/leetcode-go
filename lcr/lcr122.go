package lcr

import "bytes"

func pathEncryption(path string) string {
	buffer := bytes.NewBuffer([]byte{})

	for _, v := range path {
		if v == '.' {
			buffer.WriteString(" ")
		} else {
			buffer.WriteRune(v)
		}
	}

	return buffer.String()

}
