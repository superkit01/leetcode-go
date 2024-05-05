package week396

func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
	if len(nums) == 1 {
		return 0
	}
	n := len(nums)
	max := 0
	for i := 0; i < n; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	for i := 0; i < n; i++ {
		nums[i] = max - nums[i]
	}

	ans := 0
	//单个+1的消耗更低
	if 2*cost1 <= cost2 {
		for i := 0; i < n; i++ {
			ans = (ans + nums[i]*cost1) % 1000_000_007
		}
	}

	// TODO

	return ans
}

//最大堆
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
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
