package middle

func hIndex(citations []int) int {

	//二分
	l := 0
	r := len(citations)

	for l < r {
		mid := (l + r + 1) / 2
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}

		if cnt >= mid {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l

}
