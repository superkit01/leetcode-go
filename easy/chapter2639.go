package easy

func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))

	var calculate func(v int) int
	calculate = func(v int) int {
		if v < 0 {
			return calculate(-v) + 1
		}
		if v == 0 {
			return 1
		}
		ans := 0
		for v != 0 {
			ans += 1
			v = v / 10
		}
		return ans

	}
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		maxLength := 0
		for j := 0; j < len(grid); j++ {
			maxLength = max(maxLength, calculate(grid[j][i]))
		}
		ans[i] = maxLength
	}
	return ans
}
