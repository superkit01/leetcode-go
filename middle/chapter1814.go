package middle

func countNicePairs(nums []int) int {
	revNums := make([]int, 0)

	for _, v := range nums {
		rev := 0
		for v > 0 {
			m := v % 10
			rev = rev*10 + m
			v = v / 10
		}

		revNums = append(revNums, rev)
	}

	sub := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		sub = append(sub, nums[i]-revNums[i])
	}

	subMap := make(map[int]int, 0)

	for _, v := range sub {
		subMap[v]++
	}

	result := 0
	for _, v := range subMap {
		if v > 1 {
			result += v * (v - 1) / 2
			result = result % (1e9 + 7)
		}

	}
	return result
}
