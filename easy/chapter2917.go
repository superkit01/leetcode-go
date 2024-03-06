package easy

func findKOr(nums []int, k int) int {
	ans := 0
outer:
	for i := 0; i < 32; i++ {
		value := 1 << i
		count := 0
		for _, v := range nums {

			if v&value == value {
				count++
				if count >= k {
					ans += 1 << i
					continue outer
				}
			}

		}
	}

	return ans

}
