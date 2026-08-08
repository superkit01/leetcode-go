package lcr

import "container/heap"

func maxAltitude(heights []int, limit int) []int {
	if len(heights) == 0 || limit <= 0 {
		return []int{}
	}

	h := Heap{}
	heap.Init(&h)

	ans := make([]int, 0)

	for i := 0; i < limit-1; i++ {
		heap.Push(&h, S{heights[i], i})
	}

	for i := limit - 1; i < len(heights); i++ {
		heap.Push(&h, S{heights[i], i})
		for h.Len() > 0 {
			top := heap.Pop(&h)
			if top.(S).Index > i-limit {
				ans = append(ans, top.(S).Number)
				heap.Push(&h, top)
				break
			}
		}

	}
	return ans

}

type S struct {
	Number int
	Index  int
}

type Heap []S

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Number > h[j].Number
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(S))
}

func (h *Heap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}
