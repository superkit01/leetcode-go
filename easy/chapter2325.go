package easy

import "bytes"

func decodeMessage(key string, message string) string {
	secret := make(map[rune]rune)
	temp := 'a'
	for _, v := range key {
		if v == ' ' {
			continue
		}

		if _, ok := secret[v]; !ok {
			secret[v] = temp
			temp++
		}

	}
	result := bytes.NewBufferString("")
	for _, v := range message {
		if v == ' ' {
			result.WriteString(" ")
		} else {
			result.WriteRune(secret[v])
		}

	}
	return result.String()

}
