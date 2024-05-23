package middle

func longestEqualSubarray(nums []int, k int) int {
	group := make(map[int][]int, 0)
	for i, v := range nums {
		group[v] = append(group[v], i)
	}
	// [1,3,2,3,1,3]
	// 1:[0 4] 2:[2] 3:[1 3 5]
	ans := 0
	for _, v := range group {
		j := 0
		//[1 3 5]
		// j i
		for i := 0; i < len(v); i++ {
			for (v[i]-v[j]+1)-(i-j+1) > k {
				j++
			}
			ans = max(ans, i-j+1)
		}
	}
	return ans

}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}
