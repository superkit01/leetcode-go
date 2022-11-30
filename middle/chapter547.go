package middle

func findCircleNum(isConnected [][]int) int {
	cities := make([]bool, len(isConnected))
	province := 0
	for i := 0; i < len(isConnected); i++ {

		if !cities[i] {
			cities[i] = true
			province++
			dfs1(i, isConnected, cities)

		}

	}
	return province
}

func dfs1(i int, isConnected [][]int, cities []bool) {
	for j := 0; j < len(cities); j++ {
		if isConnected[i][j] == 1 && !cities[j] {
			cities[j] = true
			dfs1(j, isConnected, cities)
		}
	}
}
