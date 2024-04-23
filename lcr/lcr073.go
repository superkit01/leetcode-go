package lcr

//二分答案

func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := 1000000000

	for l < r {
		mid := (l + r) / 2

		hours := 0
		for i := 0; i < len(piles); i++ {
			if piles[i]%mid == 0 {
				hours += piles[i] / mid
			} else {
				hours += piles[i]/mid + 1
			}
		}

		if hours > h {
			l = mid + 1
		} else {
			r = mid
		}

	}
	return l

}
