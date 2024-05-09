package lcr

// -1 -2 3 2     ~  1

//前缀和 + Hash
func subarraySum(nums []int, k int) int {
	ans := 0

	//前缀和
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}

	prefixMap := make(map[int]int, 0)
	prefixMap[0] = 1

	for j := 0; j < len(nums); j++ {
		value := nums[j] - k
		if v, ok := prefixMap[value]; ok {
			ans += v
		}
		prefixMap[nums[j]]++
	}
	return ans

}
