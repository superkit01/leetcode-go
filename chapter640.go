package main

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	arr := strings.Split(equation, "=")

	left := ""
	for _, v := range arr[0] {
		if v == '-' {
			left += "+-"
		} else {
			left += string(v)
		}
	}
	right := ""
	for _, v := range arr[1] {
		if v == '-' {
			right += "+-"
		} else {
			right += string(v)
		}
	}

	leftArr := strings.Split(left, "+")
	rightArr := strings.Split(right, "+")

	radio := 0
	value := 0

	for i := 0; i < len(leftArr); i++ {
		if (len(leftArr[i])) == 0 {
			continue
		}
		if leftArr[i][0] == '-' {
			if leftArr[i][len(leftArr[i])-1] == 'x' {
				if len(leftArr[i]) == 2 {
					radio -= 1
				} else {
					temp, _ := strconv.Atoi(leftArr[i][1 : len(leftArr[i])-1])
					radio -= temp
				}

			} else {
				temp, _ := strconv.Atoi(leftArr[i][1:])
				value -= temp
			}

		} else {
			if leftArr[i][len(leftArr[i])-1] == 'x' {
				if len(leftArr[i]) == 1 {
					radio += 1
				} else {
					temp, _ := strconv.Atoi(leftArr[i][0 : len(leftArr[i])-1])
					radio += temp
				}
			} else {
				temp, _ := strconv.Atoi(leftArr[i])
				value += temp
			}
		}

	}

	for i := 0; i < len(rightArr); i++ {
		if (len(rightArr[i])) == 0 {
			continue
		}
		if rightArr[i][0] == '-' {
			if rightArr[i][len(rightArr[i])-1] == 'x' {
				if len(rightArr[i]) == 2 {
					radio += 1
				} else {
					temp, _ := strconv.Atoi(rightArr[i][1 : len(rightArr[i])-1])
					radio += temp
				}

			} else {
				temp, _ := strconv.Atoi(rightArr[i][1:])
				value += temp
			}

		} else {
			if rightArr[i][len(rightArr[i])-1] == 'x' {
				if len(rightArr[i]) == 1 {
					radio -= 1
				} else {
					temp, _ := strconv.Atoi(rightArr[i][0 : len(rightArr[i])-1])
					radio -= temp
				}
			} else {
				temp, _ := strconv.Atoi(rightArr[i])
				value -= temp
			}
		}

	}

	if radio == 0 && value != 0 {
		return "No solution"
	} else if radio == 0 && value == 0 {
		return "Infinite solutions"
	} else {
		return "x=" + strconv.Itoa(-value/radio)
	}

}
