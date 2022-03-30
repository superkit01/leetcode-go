package main

//import "fmt"
//
//func main() {
//	fmt.Println(lengthOfLastWord("Hello World"))
//}

func lengthOfLastWord(s string) int {
	flag := false
	start := 0
	end := 0
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		if !flag && s[i] == ' ' {
			continue
		}
		if !flag {
			start = i
			end = i
			flag = true
			continue
		}
		if flag && s[i] == ' ' {
			break
		}
		end = i

	}
	return start - end + 1
}
