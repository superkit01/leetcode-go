package lcr

type Checkout struct {
	queue          []int
	monotonicStack []int
}

func Constructor184() Checkout {
	return Checkout{
		queue:          make([]int, 0),
		monotonicStack: make([]int, 0),
	}
}

func (this *Checkout) Get_max() int {
	if len(this.monotonicStack) == 0 {
		return -1
	}
	return this.monotonicStack[0]
}

func (this *Checkout) Add(value int) {

	for len(this.monotonicStack) > 0 && this.monotonicStack[len(this.monotonicStack)-1] < value {
		this.monotonicStack = this.monotonicStack[:len(this.monotonicStack)-1]
	}
	this.monotonicStack = append(this.monotonicStack, value)
	this.queue = append(this.queue, value)
}

func (this *Checkout) Remove() int {
	if len(this.queue) == 0 {
		return -1
	}
	v := this.queue[0]
	this.queue = this.queue[1:]

	if this.monotonicStack[0] == v {
		this.monotonicStack = this.monotonicStack[1:]
	}
	return v
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
