package week336

func BeautifulSubarrays(nums []int) int64 {
	if len(nums) == 1 {
		if nums[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	prefix := make([]int, len(nums))
	prefixMap := make(map[int]int, 0)

	prefix[0] = nums[0]
	prefixMap[prefix[0]] += 1

	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i] ^ prefix[i-1]
		prefixMap[prefix[i]] += 1
	}

	var result int64
	for k, v := range prefixMap {
		if k == 0 {
			v += 1
		}

		if v == 1 {
			continue
		}
		result += int64(v) * int64((v - 1)) / 2

	}
	return result

}
