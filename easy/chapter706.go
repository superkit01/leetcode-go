package easy

type MyHashMap struct {
	V []int
	B []bool
}

func Constructor4() MyHashMap {

	return MyHashMap{
		V: make([]int, 10000001),
		B: make([]bool, 10000001),
	}

}

func (this *MyHashMap) Put(key int, value int) {
	this.B[key] = true
	this.V[key] = value

}

func (this *MyHashMap) Get(key int) int {
	if !this.B[key] {
		return -1
	}
	return this.V[key]

}

func (this *MyHashMap) Remove(key int) {
	this.B[key] = false
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
