package middle

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i int, j int) bool {
		return points[i][0] > points[j][0]
	})

	max := 0

	for i := 1; i < len(points); i++ {
		temp := points[i][0] - points[i-1][0]
		if max < temp {
			max = temp
		}
	}

	return max
}
