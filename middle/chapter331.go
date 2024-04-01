package middle

import "strings"

//  x## ==> # 

func isValidSerialization(preorder string) bool {
	chars := strings.Split(preorder, ",")

	stack := make([]string, 0)
outer:
	for _, v := range chars {

		stack = append(stack, v)

		for len(stack) >= 3 {
			if stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
				stack = stack[0 : len(stack)-3]
				stack = append(stack, "#")
			} else {
				continue outer
			}
		}
	}

	return len(stack) == 1 && stack[0] == "#"

}
