package easy

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	cache := make(map[int][]int, 0)
	for _, v := range edges {
		if _, ok := cache[v[0]]; !ok {
			cache[v[0]] = make([]int, 0)
		}
		cache[v[0]] = append(cache[v[0]], v[1])

		if _, ok := cache[v[1]]; !ok {
			cache[v[1]] = make([]int, 0)
		}
		cache[v[1]] = append(cache[v[1]], v[0])
	}

	boolMap := make(map[int]bool, 0)

	queue := make([]int, 0)
	queue = append(queue, source)
	boolMap[source] = true

	for len(queue) > 0 {
		tempQueue := make([]int, 0)
		for _, v := range queue {
			values := cache[v]
			for _, e := range values {
				if e == destination {
					return true
				}

				if !boolMap[e] {
					tempQueue = append(tempQueue, e)
					boolMap[e] = true
				}
			}

		}
		queue = tempQueue
	}
	return false

}
