package middle

type DataStream struct {
	queue   []int
	value   int
	k       int
	current int
	last    int
}

func Constructor1(value int, k int) DataStream {
	return DataStream{
		make([]int, 0),
		value,
		k,
		-1,
		-1,
	}

}

func (this *DataStream) Consec(num int) bool {
	this.queue = append(this.queue, num)
	this.current++

	if num != this.value {
		this.last = this.current
	}

	if len(this.queue) < this.k {
		return false
	}

	if len(this.queue) > this.k {
		this.queue = this.queue[1:]
	}

	return this.current-this.last >= this.k
}
