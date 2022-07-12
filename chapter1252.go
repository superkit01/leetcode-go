package main

func oddCells(m int, n int, indices [][]int) int {
	var row []int = make([]int, m)
	var col []int = make([]int, n)
	for _, e := range indices {
		row[e[0]] += 1
		col[e[1]] += 1
	}

	oddNum := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (row[i]+col[j])%2 == 1 {
				oddNum += 1
			}
		}
	}

	return oddNum

}
