package lcr

func validateBookSequences(putIn []int, takeOut []int) bool {
	stack := make([]int, 0)
	i := 0
outer:
	for _, v := range putIn {
		stack = append(stack, v)
		for len(stack) > 0 {
			if stack[len(stack)-1] == takeOut[i] {
				i++
				stack = stack[0 : len(stack)-1]
			} else {
				continue outer
			}
		}
	}

	return i == len(takeOut)
}
