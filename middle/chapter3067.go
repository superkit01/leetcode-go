package middle

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	graph := make(map[int][][2]int, 0)
	for i := 0; i <= len(edges); i++ {
		graph[i] = make([][2]int, 0)
	}
	//建图
	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], [2]int{v[1], v[2]})
		graph[v[1]] = append(graph[v[1]], [2]int{v[0], v[2]})
	}

	//dfs    current作为根节点下，可连接的节点数目
	var dfs func(parent int, current int, mod int) int
	dfs = func(parent int, current int, mod int) int {
		cnt := 0
		if mod == 0 {
			cnt++
		}
		for _, v := range graph[current] {
			if v[0] == parent {
				continue
			}
			cnt += dfs(current, v[0], (v[1]+mod)%signalSpeed)
		}
		return cnt
	}

	ans := make([]int, len(edges)+1)
	//枚举i作为根节点
	for i := 0; i <= len(edges); i++ {
		pre := 0
		for _, v := range graph[i] {
			//i的子树v的可连接节点数量
			cnt := dfs(i, v[0], v[1]%signalSpeed)
			ans[i] += pre * cnt
			pre += cnt
		}
	}
	return ans
}
