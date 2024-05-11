package lcr

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	min := math.MaxInt32
	for _, v := range triangle[len(triangle)-1] {
		if min > v {
			min = v
		}

	}
	return min
}
