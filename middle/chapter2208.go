package middle

import "container/heap"

type Item struct {
	Value float64
}

type ItemHeap []Item

func (h ItemHeap) Len() int { return len(h) }

func (h ItemHeap) Less(i, j int) bool {
	return h[i].Value > h[j].Value
}

func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(val interface{}) {
	*h = append(*h, val.(Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HalveArray(nums []int) int {
	sum := 0
	itemHeap := &ItemHeap{}
	for _, e := range nums {
		sum += e
		heap.Push(itemHeap, Item{
			Value: float64(e),
		})
	}

	var target float64
	count := 0

	for target < float64(sum)/2 {
		val := heap.Pop(itemHeap).(Item).Value
		target += val / 2
		heap.Push(itemHeap, Item{
			Value: val / 2,
		})
		count += 1
	}
	return count

}
