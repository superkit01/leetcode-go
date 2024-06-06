package hard

import (
	"leetcode-go/template"
	"math"
	"sort"
)

func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)

	unionFind := template.New(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] == 1 {
				unionFind.Union(i, j)
			}
		}
	}

	connectedConnect := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		p := unionFind.Find(i)
		connectedConnect[p]++
	}

	infected := make(map[int][]int, 0)
	for _, v := range initial {
		p := unionFind.Find(v)
		if _, ok := infected[p]; !ok {
			infected[p] = make([]int, 0)
		}
		infected[p] = append(infected[p], v)
	}

	count := 0
	ans := -1
	for k, v := range infected {
		if len(v) == 1 {

			if count < connectedConnect[k] {
				count = connectedConnect[k]
				ans = v[0]
			} else if count == connectedConnect[k] {
				//多个答案返回最小的节点
				ans = int(math.Min(float64(ans), float64(v[0])))
			}
		}
	}

	if ans == -1 {
		sort.Ints(initial)
		return initial[0]
	}
	return ans

}

// type UnionFind struct {
// 	parent []int
// }

// func New(n int) *UnionFind {
// 	parent := make([]int, n)

// 	for i := range parent {
// 		parent[i] = i
// 	}

// 	return &UnionFind{
// 		parent: parent,
// 	}
// }

// func (this *UnionFind) Find(p int) int {
// 	if this.parent[p] == p {
// 		return this.parent[p]
// 	}

// 	// return this.Find(this.parent)
// 	// path compression  路径压缩
// 	this.parent[p] = this.Find(this.parent[p])
// 	return this.parent[p]

// }

// func (this *UnionFind) Union(p, q int) {
// 	pParent := this.Find(p)
// 	qParent := this.Find(q)
// 	if pParent == qParent {
// 		return
// 	}
// 	this.parent[qParent] = pParent

// }
