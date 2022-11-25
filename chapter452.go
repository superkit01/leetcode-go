package main

import (
	"math"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	mergedSlice := make([][]int, 0)

	current := points[0]
	for i := 1; i < len(points); i++ {
		if current[0] > points[i][1] || current[1] < points[i][0] {
			mergedSlice = append(mergedSlice, current)
			current = points[i]
		} else {
			current[0] = int(math.Max(float64(points[i][0]), float64(current[0])))
			current[1] = int(math.Min(float64(points[i][1]), float64(current[1])))
		}

	}

	mergedSlice = append(mergedSlice, current)

	return len(mergedSlice)

}

func main() {

	findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})
}
