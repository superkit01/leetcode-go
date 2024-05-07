package hard

func cherryPickup(grid [][]int) int {

	//记忆化搜索
	cache := make([][][]int, len(grid))
	for deep := range cache {
		cache[deep] = make([][]int, len(grid[0]))
		for pos1 := range cache[deep] {
			cache[deep][pos1] = make([]int, len(grid[0]))
			for pos2 := range cache[deep][pos1] {
				cache[deep][pos1][pos2] = -1
			}
		}
	}

	return dfs4(grid, 0, 0, len(grid[0])-1, &cache)
}

// 深度为deep ,两个机器人位置分别为pos1和pos2的 最大摘取樱桃数量
func dfs4(grid [][]int, deep int, pos1, pos2 int, cache *[][][]int) int {
	if (*cache)[deep][pos1][pos2] != -1 {
		return (*cache)[deep][pos1][pos2]
	}

	//Break Case
	if deep == len(grid)-1 {
		if pos1 == pos2 {
			return grid[deep][pos1]
		}
		(*cache)[deep][pos1][pos2] = grid[deep][pos1] + grid[deep][pos2]
		return grid[deep][pos1] + grid[deep][pos2]
	}

	value := -1

	// value = max(dfs4(grid, deep+1, pos1-1, pos2-1), value)
	// value = max(dfs4(grid, deep+1, pos1-1, pos2), value)
	// value = max(dfs4(grid, deep+1, pos1-1, pos2+1), value)
	// value = max(dfs4(grid, deep+1, pos1, pos2-1), value)
	// value = max(dfs4(grid, deep+1, pos1, pos2), value)
	// value = max(dfs4(grid, deep+1, pos1, pos2+1), value)
	// value = max(dfs4(grid, deep+1, pos1+1, pos2-1), value)
	// value = max(dfs4(grid, deep+1, pos1+1, pos2), value)
	// value = max(dfs4(grid, deep+1, pos1+1, pos2+1), value)

	if pos1-1 >= 0 {
		if pos2-1 >= 0 {
			value = max(dfs4(grid, deep+1, pos1-1, pos2-1, cache), value)
		}
		value = max(dfs4(grid, deep+1, pos1-1, pos2, cache), value)
		if pos2+1 < len(grid[0]) {
			value = max(dfs4(grid, deep+1, pos1-1, pos2+1, cache), value)
		}
	}
	if pos2-1 >= 0 {
		value = max(dfs4(grid, deep+1, pos1, pos2-1, cache), value)
	}
	value = max(dfs4(grid, deep+1, pos1, pos2, cache), value)
	if pos2+1 < len(grid[0]) {
		value = max(dfs4(grid, deep+1, pos1, pos2+1, cache), value)
	}

	if pos1+1 < len(grid[0]) {
		if pos2-1 >= 0 {
			value = max(dfs4(grid, deep+1, pos1+1, pos2-1, cache), value)
		}
		value = max(dfs4(grid, deep+1, pos1+1, pos2, cache), value)
		if pos2+1 < len(grid[0]) {
			value = max(dfs4(grid, deep+1, pos1+1, pos2+1, cache), value)
		}
	}

	if pos1 == pos2 {
		(*cache)[deep][pos1][pos2] = grid[deep][pos1] + value
		return grid[deep][pos1] + value
	} else {
		(*cache)[deep][pos1][pos2] = grid[deep][pos1] + grid[deep][pos2] + value
		return grid[deep][pos1] + grid[deep][pos2] + value
	}

}
