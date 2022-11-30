package easy

import "math"

//balloon
func maxNumberOfBalloons(text string) int {
	container := make(map[int32]int32, 0)

	for _, v := range text {
		if _, ok := container[v]; ok {
			container[v] = container[v] + 1
		} else {
			container[v] = 1
		}
	}

	if _, ok := container['b']; !ok {
		return 0
	}
	if _, ok := container['a']; !ok {
		return 0
	}
	if _, ok := container['l']; !ok {
		return 0
	}
	if _, ok := container['o']; !ok {
		return 0
	}
	if _, ok := container['n']; !ok {
		return 0
	}

	return int(math.Min(float64(container['b']),
		math.Min(float64(container['a']),
			math.Min(float64(container['l']/2),
				math.Min(float64(container['o']/2),
					float64(container['n']))))))

}
