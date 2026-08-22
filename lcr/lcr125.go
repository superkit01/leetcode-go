package lcr

type CQueue struct {
	popStack  []int
	pushStack []int
}

func ConstructorI() CQueue {
	return CQueue{
		popStack:  make([]int, 0),
		pushStack: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.pushStack = append(this.pushStack, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.popStack) == 0 {
		for len(this.pushStack) > 0 {
			top := this.pushStack[len(this.pushStack)-1]
			this.popStack = append(this.popStack, top)
			this.pushStack = this.pushStack[0 : len(this.pushStack)-1]
		}
	}

	if len(this.popStack) == 0 {
		return -1
	} else {
		top := this.popStack[len(this.popStack)-1]
		this.popStack = this.popStack[0 : len(this.popStack)-1]
		return top
	}

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
