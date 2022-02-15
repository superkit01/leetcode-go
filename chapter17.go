package main

import "fmt"



func letterCombinations(digits string) []string {

	if len(digits)==0{
		return make([]string,0)
	}
	dict := make(map[int32][]string)
	dict['2'] = []string{"a", "b", "c"}
	dict['3'] = []string{"d", "e", "f"}
	dict['4'] = []string{"g", "h", "i"}
	dict['5'] = []string{"j", "k", "l"}
	dict['6'] = []string{"m", "n", "o"}
	dict['7'] = []string{"p", "q", "r", "s"}
	dict['8'] = []string{"t", "u", "v"}
	dict['9'] = []string{"w", "x", "y", "z"}

	tmp := make([]string, 1)

	for _, v := range digits {
		if _, ok := dict[v]; ok {
			tmp = Recursive(tmp, dict[v])
		}
	}
	return tmp
}

func Recursive(init []string, tmp []string) []string {
	result := make([]string, 0)
	for _, v := range init {
		for _, n := range tmp {
			result = append(result, fmt.Sprintf("%v%v", v, n))
		}
	}
	return result
}
