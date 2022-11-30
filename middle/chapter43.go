package middle

import (
	"fmt"
)

//func main() {
//	fmt.Println(multiply("9", "9"))
//}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmpAdd := make([]string, 0)

	for i, v := range num2 {
		tmp := multiplyTmp(num1, int(v), len(num2)-1-i)
		tmpAdd = append(tmpAdd, tmp)
	}

	result := ""
	for i := 0; i < len(tmpAdd); i++ {
		result = add(result, tmpAdd[i])
	}
	return result
}

func multiplyTmp(num1 string, v int, offset int) string {
	i := len(num1) - 1
	m := v - '0'
	result := ""
	add := 0
	for i >= 0 || add != 0 {
		j := 0
		if i >= 0 {
			j = int(num1[i]) - '0'
		}
		tmp := j*m + add
		tmpNum := tmp % 10
		add = tmp / 10
		result = fmt.Sprintf("%v%v", tmpNum, result)
		i--
	}

	for k := 0; k < offset; k++ {
		result = fmt.Sprintf("%v%v", result, "0")
	}
	return result

}

func add(num1 string, num2 string) string {
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
