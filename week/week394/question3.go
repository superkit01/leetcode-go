package week394

import (
	"fmt"
	"math"
)

func MinimumOperations(grid [][]int) int {
	n := len(grid[0])
	dp := make([][10]int, n) //将 第i列 全部变为0-9的花费的次数

	// [[0,6,2],[9,0,9],[4,9,6]]
	//  变为 	 0 1 2 3 4 5 6 7 8 9
	// 列 0 	[2 3 3 3 2 3 3 3 3 2]
	// 列 1 	[2 3 3 3 3 3 2 3 3 2]
	// 列 2 	[3 3 2 3 3 3 2 3 3 2]
	for i := 0; i < n; i++ {
		for j := 0; j <= 9; j++ {
			count := 0
			for m := 0; m < len(grid); m++ {
				if grid[m][i] != j {
					count++
				}
			}
			dp[i][j] = count
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			min := math.MaxInt32
		inner:
			for m := 0; m <= 9; m++ {
				if m == j {
					continue inner
				}
				if min > dp[i-1][m]+dp[i][j] {
					min = dp[i-1][m] + dp[i][j]
				}
			}
			dp[i][j] = min
		}
	}
	fmt.Printf("%v \n", dp)

	min := math.MaxInt32

	for i := 0; i < len(dp[len(dp)-1]); i++ {
		if dp[len(dp)-1][i] < min {
			min = dp[len(dp)-1][i]
		}
	}
	return min
}
