package easy

func finalString(s string) string {
	stack := make([]rune, 0)
	reverseStack := make([]rune, 0)

	for _, v := range s {
		if v == 'i' {
			for len(stack) > 0 {
				last := stack[len(stack)-1]
				reverseStack = append(reverseStack, last)
				stack = stack[0 : len(stack)-1]
			}

			stack = reverseStack
			reverseStack = make([]rune, 0)
		} else {
			stack = append(stack, v)
		}
	}

	return string(stack)

}
