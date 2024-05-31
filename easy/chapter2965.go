package easy

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	cnt := make([]bool, n*n)

	ans := make([]int, 2)

	for i := range grid {
		for j := range grid[i] {
			if cnt[grid[i][j]-1] == true {
				ans[0] = grid[i][j]
			}
			cnt[grid[i][j]-1] = true
		}
	}

	for i, v := range cnt {
		if !v {
			ans[1] = i + 1
			break
		}
	}

	return ans

}
