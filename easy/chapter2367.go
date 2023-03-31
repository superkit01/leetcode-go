package easy

func ArithmeticTriplets(nums []int, diff int) int {
	cache := make(map[int]bool, 0)
	result := 0

	for _, v := range nums {
		cache[v] = true
	}

	for i := 0; i < len(nums)-2; i++ {
		if _, ok := cache[nums[i]+diff]; ok {
			if _, ok2 := cache[nums[i]+diff+diff]; ok2 {
				result += 1
			}
		}
	}
	return result

}
