package easy

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	} else if len(t) < len(s) {
		return false
	}

	i := 0
	j := 0

	for j < len(t) {

		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}

		if i >= len(s) {
			return true
		}
	}

	return false

}

// func main() {
// 	isSubsequence("abc", "ahbgdc")
// }
