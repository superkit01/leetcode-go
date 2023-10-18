package middle

import (
	"container/heap"
	"math"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	h := &Heap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	var result int64

	for i := 0; i < k; i++ {
		v := heap.Pop(h).(int)
		result += int64(v)
		heap.Push(h, int(math.Ceil(float64(v)/3)))
	}

	return result
}
