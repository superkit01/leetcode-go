package easy

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}

	sum := 0
	for _, v := range code {
		sum += v
	}

	if k > 0 {

		for i := 1; i <= k; i++ {
			ans[0] += code[i]
		}

		for i := 1; i < n; i++ {
			ans[i] = ans[i-1] - code[i]
			ans[i] += code[(i+k)%n]
		}
	}

	if k < 0 {

		for i := k; i <= -1; i++ {
			ans[0] += code[i+n]
		}

		for i := 1; i < n; i++ {
			ans[i] = ans[i-1] + code[i-1]
			ans[i] -= code[(i-1+k+n)%n]
		}
	}
	return ans

}
