package lcr

func SingleNumber(nums []int) int {

	cache := make([]int, 32)

	for _, v := range nums {
		i := 0
		for v != 0 && i < 32 {
			if v&0b1 == 1 {
				cache[i]++
			}
			v = v >> 1
			i++
		}
	}

	ans := int32(0)
	for i, v := range cache {
		if v%3 == 1 {
			ans |= (1 << i)
		}

	}
	return int(ans)
}
