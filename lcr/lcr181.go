package lcr

import "strings"

func reverseMessage(message string) string {
	owords := strings.Split(message, " ")
	words := []string{}
	for _, v := range owords {
		if v == "" {
			continue
		}

		words = append(words, v)
	}

	i, j := 0, len(words)-1

	for i < j {

		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}
