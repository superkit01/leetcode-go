package easy

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor2() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}

}

// func (this *MyQueue) Push(x int) {
// 	for len(this.stack1) != 0 {
// 		v := this.stack1[len(this.stack1)-1]
// 		this.stack1 = this.stack1[0 : len(this.stack1)-1]
// 		this.stack2 = append(this.stack2, v)
// 	}

// 	this.stack1 = append(this.stack1, x)

// 	for len(this.stack2) != 0 {
// 		v := this.stack2[len(this.stack2)-1]
// 		this.stack2 = this.stack2[0 : len(this.stack2)-1]
// 		this.stack1 = append(this.stack1, v)
// 	}
// }

// func (this *MyQueue) Pop() int {
// 	v := this.stack1[len(this.stack1)-1]
// 	this.stack1 = this.stack1[0 : len(this.stack1)-1]
// 	return v

// }

// func (this *MyQueue) Peek() int {
// 	return this.stack1[len(this.stack1)-1]

// }

// func (this *MyQueue) Empty() bool {
// 	return len(this.stack1) == 0
// }

func (this *MyQueue) Push(x int) {
	this.stack2 = append(this.stack2, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stack1) == 0 {
		for len(this.stack2) != 0 {
			v := this.stack2[len(this.stack2)-1]
			this.stack2 = this.stack2[0 : len(this.stack2)-1]
			this.stack1 = append(this.stack1, v)
		}
	}

	v := this.stack1[len(this.stack1)-1]
	this.stack1 = this.stack1[0 : len(this.stack1)-1]
	return v

}

func (this *MyQueue) Peek() int {
	if len(this.stack1) == 0 {
		for len(this.stack2) != 0 {
			v := this.stack2[len(this.stack2)-1]
			this.stack2 = this.stack2[0 : len(this.stack2)-1]
			this.stack1 = append(this.stack1, v)
		}
	}
	return this.stack1[len(this.stack1)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}
