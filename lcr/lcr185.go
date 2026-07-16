package lcr

import (
	"math"
)

/*
*

K个骰子，求点数和为n的概率

	1.....           .6*k

1个骰子      1 1 1 1 1 1 0 0 0  ...  0
2个骰子      0 1 2 3 4 5 6 6 5 4 3 2 1
3个骰子      0 0 1 3 6 10 15 21 25 27 27 25 21 15 10 6 3 1
4个骰子      0 0 0 1 4 10 20 35 56 80 104 125 140 146 140 ...
5个骰子      f(i,j) = f(i-1,j-1) + f(i-1,j-2) + f(i-1,j-3) + f(i-1,j-4) + f(i-1,j-5) + f(i-1,j-6)  <=> f(i,j) = f(i-1,j-1) + f(i,j-1) - f(i-1,j-7)
6个骰子		 ...

*
*/
func statisticsProbability(num int) []float64 {
	dp := make([][]int, num)
	for i := 0; i < num; i++ {
		dp[i] = make([]int, 6*num)
	}
	for i := 0; i < 6; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < num; i++ {
		for j := 0; j < 6*num; j++ {
			v := 0
			for k := 1; k <= 6; k++ {
				if j-k >= 0 {
					v += dp[i-1][j-k]
				}
			}
			dp[i][j] = v
		}
	}

	res := make([]float64, 6*num-num+1)

	for i := 0; i < len(res); i++ {
		res[i] = float64(dp[num-1][num-1+i]) / float64(math.Pow(6, float64(num)))
	}
	return res

}
