package middle

import "math"

func MinPathCost(grid [][]int, moveCost [][]int) int {
	result := grid[0]

	for i := 1; i < len(grid); i++ {
		temp := make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			min := math.MaxInt
			for k := 0; k < len(grid[0]); k++ {
				min = int(math.Min(float64(result[k]+moveCost[grid[i-1][k]][j]), float64(min)))
			}
			temp[j] = min + grid[i][j]
		}
		result = temp

	}

	min := result[0]
	for i := 0; i < len(result); i++ {
		min = int(math.Min(float64(min), float64(result[i])))

	}
	return min

}
