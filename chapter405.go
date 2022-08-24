package main

var cache = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
}

func toHex(num int) string {

	if num == 0 {
		return "0"
	}

	arr := make([]string, 0)
	for i := 7; i >= 0; i-- {
		arr = append(arr, cache[(num>>(4*i))&0xf])
	}
	for len(arr) > 0 {
		if arr[0] == "0" {
			arr = arr[1:]
		} else {
			break
		}
	}

	result := ""

	for i := 0; i < len(arr); i++ {
		result = result + arr[i]
	}
	return result

}
