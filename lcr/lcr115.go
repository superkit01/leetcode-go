package lcr

func sequenceReconstruction(nums []int, sequences [][]int) bool {

	inDegree := make([]int, len(nums)+1)

	graph := map[int][]int{}
	//建图 遍历入度
	for _, v := range sequences {
		for k := 1; k < len(v); k++ {
			inDegree[v[k]]++
			graph[v[k-1]] = append(graph[v[k-1]], v[k])

		}

	}

	index := 0

	queue := []int{}
	for i, v := range inDegree {
		if i == 0 {
			continue
		}
		if v == 0 {
			queue = append(queue, i)
			inDegree[i] = -1
		}
	}

	for len(queue) > 0 {
		//有多个入度为0,答案不唯一
		if len(queue) > 1 {
			return false
		}
		index++
		for _, v := range graph[queue[0]] {
			inDegree[v]--
		}

		temp := []int{}
		for i, v := range inDegree {
			if i == 0 {
				continue
			}
			if v == 0 {
				temp = append(temp, i)
				inDegree[i] = -1
			}

		}
		queue = temp
	}

	
	if index == len(nums) {
		return true
	} else {
		return false
	}

}
