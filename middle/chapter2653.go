package middle

func getSubarrayBeauty(nums []int, k int, x int) []int {
	if len(nums) == 1 {
		if nums[0] > 0 {
			return []int{0}
		}
		return []int{nums[0]}
	}

	ans := make([]int, len(nums)-k+1)

	//值域很少 统计值域的个数
	count := make([]int, 101)
	for i := 0; i < k; i++ {
		count[nums[i]+50]++
	}

	xMin := func(count []int, x int) int {
		xCount := 0
		result := 0
		for i := 0; i < len(count); i++ {
			xCount += count[i]
			if xCount >= x {
				if i-50 < 0 {
					result = i - 50
				} else {
					result = 0
				}
				break
			}
		}
		return result
	}

	ans[0] = xMin(count, x)

	for i := 1; i < len(nums)-k+1; i++ {
		count[nums[i-1]+50]--
		count[nums[i+k-1]+50]++
		ans[i] = xMin(count, x)
	}
	return ans

}
