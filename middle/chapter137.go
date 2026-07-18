package middle

func SingleNumber3(nums []int) int {
	var ans int32 = 0
	for i := 0; i < 32; i++ {
		count := 0

		for _, v := range nums {
			count += v >> i & 1
		}
		if count%3 == 1 {
			ans |= 1 << i
		}

	}
	return int(ans)
}
