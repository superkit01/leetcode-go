package lcr

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n == 1 {
		return 1
	}

	uf := New(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	cache := make(map[int]bool, 0)

	for i := 0; i < n; i++ {
		cache[uf.Find(i)] = true
	}
	return len(cache)

}

type UnionFind struct {
	parent []int
}

func (uf *UnionFind) Find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFind) Union(p, q int) {
	pParent := uf.Find(p)
	qParent := uf.Find(q)
	if pParent == qParent {
		return
	}
	uf.parent[pParent] = qParent
}

func New(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}
