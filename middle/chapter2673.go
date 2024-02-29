package middle

func minIncrements(n int, cost []int) int {

	ans := 0

	for i := n - 2; i > 0; i -= 2 {
		if cost[i] > cost[i+1] {
			cost[i], cost[i+1] = cost[i+1], cost[i]
		}
		ans += cost[i+1] - cost[i]
		cost[i/2] = cost[i/2] + cost[i+1]
	}

	return ans
}
