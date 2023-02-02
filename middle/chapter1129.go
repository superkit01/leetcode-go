package middle

import "math"

func shortestAlternatingPaths2(n int, redEdges [][]int, blueEdges [][]int) []int {
	redStart := make([]int, n)
	for i := 0; i < n; i++ {
		redStart[i] = -1
	}

	dealRedStart(&redStart, redEdges, blueEdges)

	blueStart := make([]int, n)
	for i := 0; i < n; i++ {
		blueStart[i] = -1
	}

	dealRedStart(&blueStart, blueEdges, redEdges)

	result := make([]int, 0)
	result = append(result, 0)

	for i := 1; i < n; i++ {
		if redStart[i] == -1 {
			result = append(result, blueStart[i])
		} else if blueStart[i] == -1 {
			result = append(result, redStart[i])
		} else {
			result = append(result, int(math.Min(float64(redStart[i]), float64(blueStart[i]))))
		}
	}

	return result

}

func dealRedStart(redStart *[]int, redEdges [][]int, blueEdges [][]int) {
	queue := make([]int, 0)
	queue = append(queue, 0)

	redCache := make(map[int]bool, 0)
	redCache[0] = true

	blueCache := make(map[int]bool, 0)

	toggle := true
	depth := 0
	for len(queue) > 0 {
		depth++
		temp := make([]int, 0)
		if toggle {
			for _, v := range queue {
				for _, red := range redEdges {
					if red[0] == v {
						if (*redStart)[red[1]] == -1 {
							(*redStart)[red[1]] = depth
						}

						if _, ok := redCache[red[1]]; !ok {
							temp = append(temp, red[1])
							redCache[red[1]] = true

						}
					}
				}
			}
		} else {
			for _, v := range queue {
				for _, blue := range blueEdges {
					if blue[0] == v {
						if (*redStart)[blue[1]] == -1 {
							(*redStart)[blue[1]] = depth
						}

						if _, ok := blueCache[blue[1]]; !ok {
							temp = append(temp, blue[1])
							blueCache[blue[1]] = true
						}
					}
				}
			}
		}

		toggle = !toggle
		queue = temp
	}
}

/**-------------------------------------------------------------*/

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redStart := make([]int, 0)
	redStart = append(redStart, 0)
	for i := 1; i < n; i++ {
		redStart = append(redStart, findSteps(i, redEdges, blueEdges))
	}

	blueStart := make([]int, 0)
	blueStart = append(blueStart, 0)
	for i := 1; i < n; i++ {
		blueStart = append(blueStart, findSteps(i, blueEdges, redEdges))
	}

	result := make([]int, 0)
	result = append(result, 0)

	for i := 1; i < n; i++ {
		if redStart[i] == -1 {
			result = append(result, blueStart[i])
		} else if blueStart[i] == -1 {
			result = append(result, redStart[i])
		} else {
			result = append(result, int(math.Min(float64(redStart[i]), float64(blueStart[i]))))
		}
	}

	return result
}

func findSteps(n int, redEdges [][]int, blueEdges [][]int) int {

	queue := make([]int, 0)
	queue = append(queue, 0)

	redCache := make(map[int]bool, 0)
	redCache[0] = true

	blueCache := make(map[int]bool, 0)

	toggle := true
	depth := 0
	for len(queue) > 0 {
		depth++
		temp := make([]int, 0)
		if toggle {
			for _, v := range queue {
				for _, red := range redEdges {
					if red[0] == v {
						if red[1] == n {
							return depth
						}
						if _, ok := redCache[red[1]]; !ok {
							temp = append(temp, red[1])
							redCache[red[1]] = true
						}
					}
				}
			}
		} else {
			for _, v := range queue {
				for _, blue := range blueEdges {
					if blue[0] == v {
						if blue[1] == n {
							return depth
						}
						if _, ok := blueCache[blue[1]]; !ok {
							temp = append(temp, blue[1])
							blueCache[blue[1]] = true
						}
					}
				}
			}
		}

		toggle = !toggle
		queue = temp
	}

	return -1

}
