package middle

func longestConsecutive(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	numsMap := make(map[int]bool, 0)
	for _, v := range nums {
		numsMap[v] = true
	}

	ans := 0

	for k, _ := range numsMap {
		//先找到递增值最大的值
		if _, ok := numsMap[k+1]; ok {
			continue
		}

		currentValue := k - 1

		for {
			if _, ok := numsMap[currentValue]; ok {
				currentValue--
			} else {
				break
			}
		}
		if ans < k-currentValue {
			ans = k - currentValue
		}
	}

	return ans

}
