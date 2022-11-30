package easy

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	target := s + s
	return strings.Contains(target, goal)

}
