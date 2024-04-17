package topics

type UnionFind struct {
	n      int
	parent []int
}

// 并查集
func (this *UnionFind) Find(p int) int {
	if p != this.parent[p] {
		// return this.Find(this.parent[p])

		// path compression  路径压缩
		this.parent[p] = this.Find(this.parent[p])
	}
	return this.parent[p]
}

func (this *UnionFind) Union(p, q int) {
	pRoot := this.Find(p)
	qRoot := this.Find(q)
	if pRoot == qRoot {
		return
	}
	this.parent[pRoot] = qRoot
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
