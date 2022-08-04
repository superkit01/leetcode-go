package main

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	i := 1
	for {
		temp := heap.Pop(h).(int)

		if i == n {
			return temp
		}

		if _, has := seen[temp*2]; !has {
			heap.Push(h, temp*2)
			seen[temp*2] = struct{}{}
		}

		if _, has := seen[temp*3]; !has {
			heap.Push(h, temp*3)
			seen[temp*3] = struct{}{}
		}

		if _, has := seen[temp*5]; !has {
			heap.Push(h, temp*5)
			seen[temp*5] = struct{}{}
		}

		i++

	}

}
