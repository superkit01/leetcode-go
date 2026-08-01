package lcr

func CountTarget(scores []int, target int) int {

	//小于target的最大值
	bisearchLower := func(scores []int, target int) int {
		l, r := 0, len(scores)-1

		for l <= r {
			mid := l + (r-l)/2
			if scores[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}

		return l-1
	}

	return bisearchLower(scores, target+1) - bisearchLower(scores, target) 

}
