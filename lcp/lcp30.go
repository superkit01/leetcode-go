package lcp

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func MagicTower(nums []int) int {
	sum := 1
	for _, v := range nums {
		sum += v
	}
	if sum <= 0 {
		return -1
	}

	prefixSum := 1
	ans := 0

	h := &IntHeap{}
	for _, v := range nums {
		if v < 0 {
			heap.Push(h, v)
		}
		prefixSum += v
		if prefixSum <= 0 {
			top := heap.Pop(h).(int)
			prefixSum -= top
			ans++
		}
	}
	return ans

}
