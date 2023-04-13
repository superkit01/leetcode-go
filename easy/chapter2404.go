package easy

func MostFrequentEven(nums []int) int {
	countMap := make(map[int]int, 0)

	for _, v := range nums {
		if v%2 == 0 {
			countMap[v]++
		}
	}

	if len(countMap) == 0 {
		return -1
	}
	count := 0
	result := 1000000
	for k, v := range countMap {
		if v > count {
			count = v
			result = k
		} else if v == count {
			if result > k {
				result = k
			}
		}
	}
	return result
}
