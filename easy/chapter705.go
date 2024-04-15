package easy

type MyHashSet struct {
	B [][]bool
}

func Constructor5() MyHashSet {
	return MyHashSet{
		B: make([][]bool, 1001),
	}

}

func (this *MyHashSet) Add(key int) {
	index := key / 1001
	mod := key % 1001
	if this.B[index] == nil {
		this.B[index] = make([]bool, 1001)
	}
	this.B[index][mod] = true

}

func (this *MyHashSet) Remove(key int) {
	index := key / 1001
	mod := key % 1001
	if this.B[index] == nil {
		return
	}
	this.B[index][mod] = false
}

func (this *MyHashSet) Contains(key int) bool {
	index := key / 1001
	mod := key % 1001
	if this.B[index] == nil {
		return false
	}
	return this.B[index][mod]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
