package topics

type UnionFind struct {
	n      int
	parent []int
}

// 并查集
func (uf UnionFind) Find(p int) int {
	if p != uf.parent[p] {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
}

func New(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		n:      n,
		parent: parent,
	}
}
