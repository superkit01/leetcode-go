package middle

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := 0
	for _, e := range rolls {
		sum += e
	}
	restSum := (m+n)*mean - sum
	if restSum < n || restSum > 6*n {
		return make([]int, 0)
	}
	result := make([]int, n)

	rest := restSum % n
	shang := restSum / n

	for i := 0; i < len(result); i++ {
		if i < rest {
			result[i] = shang + 1
		} else {
			result[i] = shang
		}
	}
	return result
}
