package main

import (
	"bytes"
	"strings"
)

func isPalindrome(s string) bool {
	var sTemp bytes.Buffer
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') {
			sTemp.WriteRune(c)
		}
	}
	s = strings.ToLower(sTemp.String())
	for i := 0; i < (len(s)+1)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true

}

//func main() {
//	isPalindrome("race a car")
//}
