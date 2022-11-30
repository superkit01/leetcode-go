package middle

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	times := len(s)/numRows + 1
	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < times; j++ {
			if j*(2*numRows-2)-i < len(s) && j*(2*numRows-2)-i >= 0 {
				res = fmt.Sprintf("%v%v", res, string(s[j*(2*numRows-2)-i]))
			}
			if i != 0 && i != numRows-1 && j*(2*numRows-2)+i < len(s) && j*(2*numRows-2)+i >= 0 {
				res = fmt.Sprintf("%v%v", res, string(s[j*(2*numRows-2)+i]))
			}
		}
	}
	return res
}
