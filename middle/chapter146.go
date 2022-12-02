package middle

type LRUCache struct {
	container map[int]*Element
	head      *Element
	tail      *Element
	size      int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		container: make(map[int]*Element),
		head: &Element{
			val: -1,
			key: -1,
		},
		tail: &Element{
			val: -1,
			key: -1,
		},
		size: capacity,
	}

	cache.head.next = cache.tail
	cache.tail.pre = cache.head
	return cache

}

type Element struct {
	pre  *Element
	next *Element
	val  int
	key  int
}

func (this *LRUCache) Get(key int) int {
	//不存在
	if _, ok := this.container[key]; !ok {
		return -1
	}

	v := this.container[key]
	this.moveToHead(v)
	return v.val

}

func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.container[key]; !ok {
		element := &Element{
			pre:  nil,
			next: nil,
			val:  value,
			key:  key,
		}
		this.container[key] = element
		this.addToHead(element)
		if len(this.container) > this.size {
			this.removeTail()
		}
	} else {
		element := this.container[key]
		element.val = value
		this.moveToHead(element)
	}

}

func (this *LRUCache) moveToHead(e *Element) {
	this.removeElement(e)
	this.addToHead(e)
}

func (this *LRUCache) removeTail() {
	e := this.tail.pre
	this.removeElement(e)
	delete(this.container, e.key)
}

func (this *LRUCache) addToHead(e *Element) {
	element := this.head.next
	e.next = element
	element.pre = e
	this.head.next = e
	e.pre = this.head
}
func (this *LRUCache) removeElement(e *Element) {
	e.pre.next = e.next
	e.next.pre = e.pre
}
