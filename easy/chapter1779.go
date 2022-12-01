package easy

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {

	minDistance := math.MaxInt
	minIndex := len(points)

	for i, v := range points {
		if v[0] == x {
			distance := int(math.Abs(float64(v[1] - y)))
			if distance < minDistance {
				minDistance = distance
				minIndex = i
			}
		}
		if v[1] == y {
			distance := int(math.Abs(float64(v[0] - x)))
			if distance < minDistance {
				minDistance = distance
				minIndex = i

			}
		}
	}

	if minIndex == len(points) {
		return -1
	} else {
		return minIndex
	}

}
