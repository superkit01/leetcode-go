package main

import (
	"math"
)

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, v := range asteroids {
		compare(&stack, v)
	}

	return stack
}

func compare(stack *[]int, v int) {
	if len(*stack) == 0 {
		*stack = append(*stack, v)
		return
	}

	top := (*stack)[len(*stack)-1]

	if top*v > 0 || top < 0 {
		*stack = append(*stack, v)
		return
	}

	if top > 0 && v < 0 {

		if math.Abs(float64(top)) > math.Abs(float64(v)) {
			return
		}
		if math.Abs(float64(top)) == math.Abs(float64(v)) {
			*stack = (*stack)[0 : len(*stack)-1]
			return
		}
		if math.Abs(float64(top)) < math.Abs(float64(v)) {
			*stack = (*stack)[0 : len(*stack)-1]
			compare(stack, v)
		}
	} else {
		*stack = append(*stack, v)
		return
	}
}
