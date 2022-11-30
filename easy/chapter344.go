package easy

func reverseString(s []byte) {
	count := len(s)
	temp := s[0]
	for i := 0; i <= (count/2)-1; i++ {
		temp = s[i]
		s[i] = s[count-1-i]
		s[count-1-i] = temp
	}
}

// func main() {
// reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
// }
