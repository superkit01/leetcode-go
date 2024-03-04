package middle

func reachableNodes(n int, edges [][]int, restricted []int) int {

	restrictedMap := make(map[int]bool, 0)
	for _, v := range restricted {
		restrictedMap[v] = true
	}

	graphMap := make(map[int][]int, 0)

	for _, v := range edges {
		if _, ok := restrictedMap[v[0]]; ok {
			continue
		}
		if _, ok := restrictedMap[v[1]]; ok {
			continue
		}

		if _, ok := graphMap[v[0]]; !ok {
			graphMap[v[0]] = make([]int, 0)
		}
		graphMap[v[0]] = append(graphMap[v[0]], v[1])

		if _, ok := graphMap[v[1]]; !ok {
			graphMap[v[1]] = make([]int, 0)
		}
		graphMap[v[1]] = append(graphMap[v[1]], v[0])
	}

	cache := make(map[int]bool, 0)
	queue := make([]int, 0)

	cache[0] = true
	queue = append(queue, 0)

	ans := 1

	for len(queue) > 0 {
		temp := make([]int, 0)
		for _, v := range queue {
			if gv, ok := graphMap[v]; ok {
				for _, e := range gv {
					if _, ok := cache[e]; !ok {
						temp = append(temp, e)
						cache[e] = true
						ans++
					}
				}
			}
		}
		queue = temp
	}

	return ans

}
