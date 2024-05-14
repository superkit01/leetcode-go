package week397

import "math"

// 枚举以 [i,j] 为终点
// 求[i,j]坐上区域的最小值之差(不包含[i,j]单元格)
// 						
// [9,5,7,	3]    
// [8,9,6,	1]
// [6,7,14, 3]
// [2,5,3,	1]

// 二维前缀最小值
// [max,max,max,max,max]    
// [max,9,	5,	5,	3]
// [max,8,	5,  5,	1]
// [max,6,	5,	5,	1]
// [max,2,	5,	3,	1]

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}
	ans := math.MinInt
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = math.MaxInt32
				continue
			}
			matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1])
			temp := grid[i-1][j-1] - matrix[i][j]

			if ans < temp {
				ans = temp
			}
			matrix[i][j] = min(matrix[i][j], grid[i-1][j-1])

		}
	}
	return ans

}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
