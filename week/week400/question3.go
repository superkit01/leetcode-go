package week400

import (
	"container/heap"
	"sort"
)

func clearStars(s string) string {

	h := &CharPosHeap{}

	for i, v := range s {
		if v == '*' {
			heap.Pop(h)
		} else {
			heap.Push(h, CharPos{
				c:   v,
				pos: i,
			})
		}
	}

	remain := []CharPos{}

	for h.Len() > 0 {
		remain = append(remain, h.Pop().(CharPos))
	}

	sort.Slice(remain, func(i, j int) bool {
		return remain[i].pos < remain[j].pos
	})

	ans := ""
	for _, v := range remain {
		ans += string(v.c)
	}
	return ans

}

type CharPos struct {
	c   rune
	pos int
}

type CharPosHeap []CharPos

func (h CharPosHeap) Len() int { return len(h) }

func (h CharPosHeap) Less(i, j int) bool {
	if h[i].c == h[j].c {
		return h[i].pos > h[j].pos
	}
	return h[i].c < h[j].c
}

func (h CharPosHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *CharPosHeap) Push(val interface{}) {
	*h = append(*h, val.(CharPos))
}

func (h *CharPosHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
