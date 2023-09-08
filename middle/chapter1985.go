package middle

import (
	"container/heap"
)

type IntHeap []string

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	if len(h[i]) < len(h[j]) {
		return true
	} else if len(h[i]) > len(h[j]) {
		return false
	} else {
		return h[i] < h[j]
	}
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(string))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KthLargestNumber(nums []string, k int) string {

	h := &IntHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	temp := ""
	d := len(*h) - k
	for i := 0; i <= d; i++ {
		temp = heap.Pop(h).(string)
	}
	return temp

}
