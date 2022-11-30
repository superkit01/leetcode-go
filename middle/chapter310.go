package middle

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	degree := make(map[int]int, 0)
	rel := make(map[int][]int, 0)

	for _, edge := range edges {
		degree[edge[0]]++
		degree[edge[1]]++
		rel[edge[1]] = append(rel[edge[1]], edge[0])
		rel[edge[0]] = append(rel[edge[0]], edge[1])

	}

	queue := make([]int, 0)
	for k, v := range degree {
		if v == 1 {
			queue = append(queue, k)
		}
	}
	temp := make([]int, 0)
	for len(queue) > 0 {
		temp = make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			degree[queue[i]]--
			temp = append(temp, queue[i])
			for _, e := range rel[queue[i]] {
				degree[e]--
				if degree[e] == 1 {
					queue = append(queue, e)
				}
			}
		}
		queue = queue[length:]
	}
	return temp
}

//func main() {
//	findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
//}
