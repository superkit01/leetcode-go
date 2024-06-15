package lcr

//并查集
func isBipartite(graph [][]int) bool {
	if len(graph) == 1 {
		return true
	}
	n := len(graph)
	parent := make([]int, 2*n)

	for i := range parent {
		parent[i] = i
	}
	uf := &UnionFind2{
		parent: parent,
	}

	for i, v := range graph {
		for _, j := range v {
			if uf.Find(i) == uf.Find(j) {
				return false
			}
			uf.Union(i, j+n)
			uf.Union(j, i+n)
		}

	}
	return true

}

type UnionFind2 struct {
	parent []int
}

func (this *UnionFind2) Find(x int) int {
	if this.parent[x] != x {
		this.parent[x] = this.Find(this.parent[x])
	}
	return this.parent[x]
}

func (this *UnionFind2) Union(x, y int) {
	pX := this.Find(x)
	pY := this.Find(y)
	if pX == pY {
		return
	}
	this.parent[pX] = pY
}

//BFS
func IsBipartiteII(graph [][]int) bool {
	//初始值 0   红 1 蓝2
	// [[],[2,4,6],[1,4,8,9],[7,8],[1,2,8,9],[6,9],[1,5,7,8,9],[3,6,9],[2,3,4,6,9],[2,4,5,6,7,8]]
	color := make([]int, len(graph))

	for index, c := range color {
		if c == 0 {
			queue := []int{index}
			color[index] = 1
			for len(queue) > 0 {
				tempQueue := map[int]int{}

				for _, v := range queue {
					if color[v] == 1 {
						for _, m := range graph[v] {
							if color[m] == 1 {
								return false
							}

							if color[m] == 0 {
								tempQueue[m] = 0
							}
							color[m] = 2

						}
					}
					if color[v] == 2 {
						for _, m := range graph[v] {
							if color[m] == 2 {
								return false
							}
							if color[m] == 0 {
								tempQueue[m] = 0
							}
							color[m] = 1
						}
					}

				}

				queue = []int{}
				for k, _ := range tempQueue {
					queue = append(queue, k)
				}

			}

		}

	}

	return true

}

//DFS
func IsBipartiteIII(graph [][]int) bool {
	//初始值 0   红 1 蓝2
	// [[],[2,4,6],[1,4,8,9],[7,8],[1,2,8,9],[6,9],[1,5,7,8,9],[3,6,9],[2,3,4,6,9],[2,4,5,6,7,8]]
	color := make([]int, len(graph))

	//给index位置染色preColor
	var dfs func(index, preColor int) bool
	dfs = func(index, preColor int) bool {
		if color[index] != 0 {
			return color[index] == preColor
		}

		color[index] = preColor

		for _, v := range graph[index] {
			if preColor == 1 {
				if !dfs(v, 2) {
					return false
				}

			} else {
				if !dfs(v, 1) {
					return false
				}
			}

		}
		return true

	}

	for index, c := range color {
		if c == 0 {
			if !dfs(index, 1) {
				return false
			}
		}
	}

	return true

}
