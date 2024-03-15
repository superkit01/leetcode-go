package lcr

import "math"

//dp[i][j]
//[1,2,3,4,5,6,7,8]
// 	 <-j   i

func LenLongestFibSubseq(arr []int) int {
	indexMap := make(map[int]int, 0)
	for i, v := range arr {
		indexMap[v] = i
	}

	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}

	for i := 1; i < len(arr)-1; i++ {
		for j := i - 1; j >= 0; j-- {
			sum := arr[i] + arr[j]
			if v, ok := indexMap[sum]; ok {
				dp[v][i] = int(math.Max(float64(dp[i][j]+1), float64(3)))
			}
		}
	}

	ans := 0
	for _, m := range dp {
		for _, n := range m {
			if ans < n {
				ans = n
			}
		}
	}
	return ans

}
