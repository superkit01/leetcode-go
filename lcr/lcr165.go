package lcr

import (
	"golang.org/x/exp/slices"
)

func crackNumber(ciphertext int) int {
	if ciphertext < 10 {
		return 1
	}
	cipher := []int{}
	for ciphertext > 0 {
		cipher = append(cipher, ciphertext%10)
		ciphertext /= 10
	}
	slices.Reverse(cipher)

	dp := make([]int, len(cipher))

	dp[0] = 1
	dp[1] = 1
	if cipher[0]*10+cipher[1] < 26 && cipher[0] != 0 {
		dp[1] = 2
	}

	for i := 2; i < len(cipher); i++ {
		if cipher[i-1]*10+cipher[i] < 26 && cipher[i-1] != 0 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}

	}
	return dp[len(dp)-1]

}
