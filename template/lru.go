package template

type LRUNode struct {
	Key  int
	Val  int
	Pre  *LRUNode
	Next *LRUNode
}

type LRUCache struct {
	container map[int]*LRUNode
	dummyHead *LRUNode
	dummyTail *LRUNode
	size      int
}

func ConstructorLRU(capacity int) LRUCache {
	cache := LRUCache{
		container: make(map[int]*LRUNode, 0),
		dummyHead: &LRUNode{},
		dummyTail: &LRUNode{},
		size:      capacity,
	}

	cache.dummyHead.Next = cache.dummyTail
	cache.dummyTail.Pre = cache.dummyHead
	return cache

}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.container[key]; !ok {
		return -1
	}
	node := this.container[key]
	this.moveToHead(node)
	return node.Val

}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.container[key]; !ok {
		node := &LRUNode{
			Val: value,
			Key: key,
		}
		this.container[key] = node
		this.addHead(node)
		if len(this.container) > this.size {
			this.removeTail()
		}
	} else {
		v.Val = value
		this.moveToHead(v)
	}

}

func (this *LRUCache) moveToHead(node *LRUNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	head := this.dummyHead.Next
	node.Next = head
	head.Pre = node

	this.dummyHead.Next = node
	node.Pre = this.dummyHead

}

func (this *LRUCache) removeTail() {
	node := this.dummyTail.Pre
	delete(this.container, node.Key)

	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) addHead(node *LRUNode) {

	this.dummyHead.Next.Pre = node
	node.Next = this.dummyHead.Next
	node.Pre = this.dummyHead
	this.dummyHead.Next = node
}
