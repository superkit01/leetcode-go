package lcr

import "math"

func countNumbers(cnt int) []int {
	ans := make([]int, int(math.Pow10(cnt)))
	for i := 0; i < int(math.Pow10(cnt)); i++ {
		ans = append(ans, i)
	}
	return ans
}
