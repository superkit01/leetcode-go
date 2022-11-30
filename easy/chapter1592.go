package main

import "strings"

func reorderSpaces(text string) string {
	textSlice := strings.Fields(text)
	spaceCount := strings.Count(text, " ")
	if len(textSlice) == 1 {
		return textSlice[0] + strings.Repeat(" ", spaceCount)
	}
	eachCount := spaceCount / (len(textSlice) - 1)

	return strings.Join(textSlice, strings.Repeat(" ", eachCount)) + strings.Repeat(" ", spaceCount%(len(textSlice)-1))

}
