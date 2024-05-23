package lcr

func findOrder(numCourses int, prerequisites [][]int) []int {
	//有相同遍历入度

	inDegree := make([]int, numCourses)
	//建图
	graph := make(map[int][]int)
	for _, v := range prerequisites {
		inDegree[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	ans := []int{}

	for i, v := range inDegree {
		if v == 0 {
			ans = append(ans, i)
		}
	}
	i := 0
	for i < len(ans) {
		for _, k := range graph[ans[i]] {
			inDegree[k]--
			if inDegree[k] == 0 {
				ans = append(ans, k)
			}
		}
		i++
	}
	if len(ans) != numCourses {
		return []int{}
	}
	return ans

}
