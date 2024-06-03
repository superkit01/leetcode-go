package dweek130

func maxPointsInsideSquare(points [][]int, s string) int {
	l := 0
	r := len(s) - 1
outer:
	for l < r {
		mid := (l + r) / 2

		cache := make(map[byte]bool, 0)
		for index, v := range points {
			if v[0] <= mid && v[0] >= -mid && v[1] <= mid && v[1] >= -mid {
				if _, ok := cache[s[index]]; ok {
					r = mid
					continue outer
				}
				cache[s[index]] = true
			}
		}

		l = mid + 1
	}

	ans := 0
	for _, v := range points {
		if v[0] <= l-1 && v[0] >= -(l-1) && v[1] <= l-1 && v[1] >= -(l-1) {
			ans++
		}
	}
	return ans
}
