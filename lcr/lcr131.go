package lcr

import "math"

//dp [i] : 长度为i的竹子能够达到的乘积最大值
func cuttingBamboo(bamboo_len int) int {
	if bamboo_len <= 4 {
		return bamboo_len / 2 * (bamboo_len - bamboo_len/2)
	}

	dp := make([]int, bamboo_len+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 4
	for i := 5; i <= bamboo_len; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			max = int(math.Max(float64(max), float64(dp[i-j]*dp[j])))
		}
		dp[i] = max
	}

	return dp[bamboo_len]
}
