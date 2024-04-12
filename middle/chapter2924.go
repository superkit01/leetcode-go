package middle

func findChampion(n int, edges [][]int) int {

	inDegree := make([]int, n)

	for _, v := range edges {
		inDegree[v[1]]++
	}

	count := 0

	win := -1

	for i, v := range inDegree {
		if v == 0 {
			win = i
			count++
		}
	}

	if count > 1 || count == 0 {
		return -1
	} else {
		return win
	}

}
