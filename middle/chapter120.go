package middle

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		temp := make([]int, len(triangle))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				temp[j] = dp[j] + triangle[i][j]
			} else if j == i {
				temp[j] = dp[j-1] + triangle[i][j]
			} else {
				temp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + triangle[i][j]
			}
		}
		dp = temp
	}
	result := math.MaxInt

	for i := 0; i < len(dp); i++ {
		if result > dp[i] {
			result = dp[i]
		}
	}
	return result

}

// func main() {

// 	minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
// }
