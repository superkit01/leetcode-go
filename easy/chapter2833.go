package easy

import "math"

func furthestDistanceFromOrigin(moves string) int {
	countL := 0
	count_ := 0

	for _, v := range moves {
		if v == 'L' {
			countL++
		} else if v == 'R' {
			countL--
		} else {
			count_++
		}
	}

	return int(math.Abs(float64(countL))) + count_

}
