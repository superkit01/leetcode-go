package middle

import (
	"container/heap"
	"sort"
)

func TotalCost(costs []int, k int, candidates int) int64 {

	if len(costs) <= 2*candidates {
		sort.Ints(costs)
		var sum int64
		for i := 0; i < k; i++ {
			sum += int64(costs[i])
		}
		return sum
	}

	left := 0
	right := len(costs) - 1

	workerHeap := &WorkerHeap{}
	//初始化堆数据
	for i := 0; i < candidates; i++ {
		heap.Push(workerHeap, &Worker{
			index:     i,
			cost:      costs[i],
			direction: 0,
		})
		left++
		heap.Push(workerHeap, &Worker{
			index:     len(costs) - 1 - i,
			cost:      costs[len(costs)-1-i],
			direction: 1,
		})
		right--
	}

	var ans int64

	for i := 0; i < k; i++ {
		worker := heap.Pop(workerHeap).(*Worker)
		ans += int64(worker.cost)
		if worker.direction == 0 {
		inner:
			for left <= right {
				heap.Push(workerHeap, &Worker{
					index:     left,
					cost:      costs[left],
					direction: 0,
				})
				left++
				break inner
			}
		}

		if worker.direction == 1 {
		inner2:
			for left <= right {
				heap.Push(workerHeap, &Worker{
					index:     right,
					cost:      costs[right],
					direction: 1,
				})
				right--
				break inner2
			}
		}
	}

	return ans
}

type Worker struct {
	index     int
	cost      int
	direction int
}

type WorkerHeap []*Worker

func (h WorkerHeap) Len() int {
	return len(h)
}

func (h WorkerHeap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].index < h[j].index
	}

	return h[i].cost < h[j].cost
}

func (h WorkerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WorkerHeap) Push(x interface{}) {
	*h = append(*h, x.(*Worker))
}

func (h *WorkerHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}
