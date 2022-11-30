package middle

import (
	"math"
)

//func minPathSum(grid [][]int) int {
//	m := len(grid)
//	n := len(grid[0])
//
//	result := make([]int, 0)
//
//	dp3(grid, m-1, n-1, 0, &result)
//
//	minV := result[0]
//	for i := 1; i < len(result); i++ {
//		if result[i] < minV {
//			minV = result[i]
//		}
//	}
//	return minV
//}
//
//func dp3(grid [][]int, restM int, restN int, score int, result *[]int) {
//	if restM == 0 && restN == 0 {
//		score += grid[restM][restN]
//		*result = append(*result, score)
//	} else if restM == 0 {
//		dp3(grid, restM, restN-1, score+grid[restM][restN], result)
//	} else if restN == 0 {
//		dp3(grid, restM-1, restN, score+grid[restM][restN], result)
//	} else {
//		dp3(grid, restM, restN-1, score+grid[restM][restN], result)
//		dp3(grid, restM-1, restN, score+grid[restM][restN], result)
//	}
//}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

//
//func main() {
//	fmt.Println("%V", minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
//}
