package hard

import (
	"math"
	"sort"
)

func MinMalwareSpreadII(graph [][]int, initial []int) int {
	sort.Ints(initial)
	n := len(graph)
	unionFind := New(n)

	initialBool := make([]bool, n)
	for _, i := range initial {
		initialBool[i] = true
	}

	//排除污染节点之后并查集
	for i := 0; i < n; i++ {
		if initialBool[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if graph[i][j] == 1 && !initialBool[j] {
				unionFind.Union(i, j)
			}
		}
	}

	//根节点的连通数量
	connected := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		if !initialBool[i] {
			p := unionFind.Find(i)
			connected[p]++
		}
	}

	/*
		0 —— (1)    2
		|
		3

		[ 1, 1, 0, 1 ]
		[ 1, 1, 0, 0 ]
		[ 0, 0, 1, 0 ]
		[ 1, 0, 0, 1 ]
	*/

	// 污染节点连通的parent set集合（使用map替代set）
	temp := make(map[int]map[int]bool, 0)
	for _, e := range initial {
		temp[e] = make(map[int]bool, 0)
		for k, v := range graph[e] { //污染节点的连通节点
			if v == 1 && !initialBool[k] {
				pK := unionFind.Find(k)
				temp[e][pK] = true
			}
		}
	}

	//污染节点连通的parent次数   parent集合 被多个污染节点连通时，移除当前节点并不能使得连通节点不受污染
	count := make([]int, n)
	for _, v := range temp {
		for i := range v {
			count[i]++
		}
	}

	//污染节点连通的parent次数==1时，（污染节点） 连通的 （干净集合节点）的数量
	ansMap := make(map[int]int, 0)
	for k, v := range temp {
		for k1, _ := range v {
			if count[k1] == 1 {
				ansMap[k] = ansMap[k] + connected[k1]
			}
		}
	}

	ans := initial[0]
	infected := 0
	for k, v := range ansMap {
		if v > infected {
			infected = v
			ans = k
		}
		if v == infected {
			ans = int(math.Min(float64(ans), float64(k)))
		}
	}

	return ans

}
