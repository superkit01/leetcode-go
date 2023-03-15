package middle

func MaximalNetworkRank(n int, roads [][]int) int {
	if len(roads) == 0 {
		return 0
	}

	matrix := make([][]bool, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}

	count := make([]int, n)

	for _, v := range roads {
		count[v[0]]++
		count[v[1]]++
		matrix[v[0]][v[1]] = true
		matrix[v[1]][v[0]] = true
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v := count[i] + count[j]

			if matrix[i][j] {
				v--
			}
			if max < v {
				max = v
			}

		}
	}

	return max

}
