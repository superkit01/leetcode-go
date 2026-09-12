package lcr

import "container/heap"

type MedianFinder struct {
	lowHeap  *LowHeap
	highHeap *HighHeap
	lowSize  int
	highSize int
}

/** initialize your data structure here. */
func Constructor160() MedianFinder {
	return MedianFinder{
		lowHeap:  &LowHeap{},
		highHeap: &HighHeap{},
		lowSize:  0,
		highSize: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {

	if this.highSize > 0 {
		highTop := heap.Pop(this.highHeap).(int)
		heap.Push(this.highHeap, highTop)

		if num >= highTop {
			heap.Push(this.highHeap, num)
			this.highSize++
		} else {
			heap.Push(this.lowHeap, num)
			this.lowSize++
		}
	} else {
		heap.Push(this.lowHeap, num)
		this.lowSize++
	}

	for this.highSize > this.lowSize {
		highTop := heap.Pop(this.highHeap).(int)
		heap.Push(this.lowHeap, highTop)
		this.highSize--
		this.lowSize++
	}

	for this.lowSize > this.highSize+1 {
		lowTop := heap.Pop(this.lowHeap).(int)
		heap.Push(this.highHeap, lowTop)
		this.highSize++
		this.lowSize--
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if (this.lowSize+this.highSize)%2 == 0 {
		lowTop := heap.Pop(this.lowHeap)
		heap.Push(this.lowHeap, lowTop)

		highTop := heap.Pop(this.highHeap)
		heap.Push(this.highHeap, highTop)

		return float64(lowTop.(int)+highTop.(int)) / 2
	} else {
		lowTop := heap.Pop(this.lowHeap)
		heap.Push(this.lowHeap, lowTop)
		return float64(lowTop.(int))
	}

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type LowHeap []int

func (h LowHeap) Len() int           { return len(h) }
func (h LowHeap) Less(i, j int) bool { return h[i] > h[j] } // 大顶堆
func (h LowHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LowHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *LowHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type HighHeap []int

func (h HighHeap) Len() int           { return len(h) }
func (h HighHeap) Less(i, j int) bool { return h[i] < h[j] } // 大顶堆
func (h HighHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HighHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HighHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
