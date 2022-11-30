package easy

import "fmt"

//func main() {
//	fmt.Println(	addStrings("11","22"))
//}

func addStrings(num1 string, num2 string) string {
	len1 := len(num1) - 1
	len2 := len(num2) - 1
	add := 0
	var res string
	for len1 >= 0 || len2 >= 0 || add != 0 {
		x := 0
		y := 0
		if len1 >= 0 {
			x = int(num1[len1] - '0')
		}
		if len2 >= 0 {
			y = int(num2[len2] - '0')
		}

		tmp := x + y + add
		add = tmp / 10
		res = fmt.Sprintf("%v%v", tmp%10, res)

		len1--
		len2--
	}

	return res
}
