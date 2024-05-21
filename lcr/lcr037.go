package lcr

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

outer:
	for _, v := range asteroids {
		for {
			if len(stack) == 0 || v >= 0 {
				stack = append(stack, v)
				continue outer
			}

			top := stack[len(stack)-1]
			if top < 0 {
				stack = append(stack, v)
				continue outer
			}

			if top == -v {
				stack = stack[0 : len(stack)-1]
				continue outer
			}
			if top > -v {
				continue outer
			}
			if top < -v {
				stack = stack[0 : len(stack)-1]
			}
		}
	}
	return stack
}
