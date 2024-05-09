package lcr

func findMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		//前缀和
		if i > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}

	prefixMap := make(map[int]int, 0)
	prefixMap[0] = -1
	prefixMap[nums[0]] = 0
	ans := 0
	for i := 1; i < len(nums); i++ {
		if v, ok := prefixMap[nums[i]]; !ok {
			prefixMap[nums[i]] = i
		} else {
			if ans < i-v {
				ans = i - v
			}
		}

	}
	return ans

}
