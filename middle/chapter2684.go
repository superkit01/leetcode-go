package middle

// [
// 	[187,167,209,251,152,236,263,128,135],
// 	[267,249,251,285,73, 204,70, 207,74],
// 	[189,159,235,66, 84, 89, 153,111,189],
// 	[120,81, 210,7,  2,  231,92, 128,218],
// 	[193,131,244,293,284,175,226,205,245]
// ]
func maxMoves(grid [][]int) int {
	if len(grid[0]) == 1 {
		return 0
	}

	queue := make([][]int, 0)

	for i := 0; i < len(grid); i++ {
		queue = append(queue, []int{i, 0})
	}

	ans := 0
	for len(queue) > 0 {
		tempQueue := make([][]int, 0)
		cache := make([]bool, len(grid))//去重
		for _, v := range queue {
			if v[1] == len(grid[0])-1 {
				return len(grid[0]) - 1
			}
			if v[0]-1 >= 0 && grid[v[0]-1][v[1]+1] > grid[v[0]][v[1]] && !cache[v[0]-1]{
				
				tempQueue = append(tempQueue, []int{v[0] - 1, v[1] + 1})
				cache[v[0]-1]=true

			}

			if v[0]+1 < len(grid) && grid[v[0]+1][v[1]+1] > grid[v[0]][v[1]]  && !cache[v[0]+1] {
				tempQueue = append(tempQueue, []int{v[0] + 1, v[1] + 1})
				cache[v[0]+1]=true
			}

			if grid[v[0]][v[1]+1] > grid[v[0]][v[1]]  && !cache[v[0]] {
				tempQueue = append(tempQueue, []int{v[0], v[1] + 1})
				cache[v[0]]=true
			}
		}
		if len(tempQueue) == 0 {
			return ans
		}
		ans++
		queue = tempQueue

	}

	return ans

}
