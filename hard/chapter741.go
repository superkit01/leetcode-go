package hard

import "math"

//[[0,1,-1],
// [1,0,-1],
// [1,1,1]]
func CherryPickupII(grid [][]int) int {
	n := len(grid)

	//记忆化搜索
	cache := make([][][]int, 2*n-1)
	for deep := range cache {
		cache[deep] = make([][]int, n)
		for pos1 := range cache[deep] {
			cache[deep][pos1] = make([]int, n)
			for pos2 := range cache[deep][pos1] {
				cache[deep][pos1][pos2] = -1
			}
		}
	}

	return max(dfsII(grid, n, 0, 0, 0, &cache), 0)
}

// 深度（走了deep步）为deep ,两个机器人纵向位置分别为pos1和pos2的 最大摘取樱桃数量
func dfsII(grid [][]int, n int, deep int, pos1, pos2 int, cache *[][][]int) int {
	if (*cache)[deep][pos1][pos2] != -1 {
		return (*cache)[deep][pos1][pos2]
	}

	//Break Case
	if grid[pos1][deep-pos1] == -1 || grid[pos2][deep-pos2] == -1 {
		(*cache)[deep][pos1][pos2] = math.MinInt
		return math.MinInt
	}
	if deep == 2*n-2 {
		(*cache)[deep][pos1][pos2] = grid[pos1][deep-pos1]
		return grid[pos1][deep-pos1]
	}

	value := math.MinInt
	//pos1向下
	if pos1+1 < n {
		//pos2 向下
		if pos2+1 < n {
			value = max(dfsII(grid, n, deep+1, pos1+1, pos2+1, cache), value)
		}
		//pos2 向右
		if deep-pos2+1 < n {
			value = max(dfsII(grid, n, deep+1, pos1+1, pos2, cache), value)
		}
	}
	//pos1向右
	if deep-pos1+1 < n {
		//pos2 向下
		if pos2+1 < n {
			value = max(dfsII(grid, n, deep+1, pos1, pos2+1, cache), value)
		}
		//pos2 向右
		if deep-pos2+1 < n {
			value = max(dfsII(grid, n, deep+1, pos1, pos2, cache), value)
		}
	}

	if pos1 == pos2 {
		(*cache)[deep][pos1][pos2] = grid[pos1][deep-pos1] + value
		return grid[pos1][deep-pos1] + value
	} else {
		(*cache)[deep][pos1][pos2] = grid[pos1][deep-pos1] + grid[pos2][deep-pos2] + value
		return grid[pos1][deep-pos1] + grid[pos2][deep-pos2] + value
	}

}
