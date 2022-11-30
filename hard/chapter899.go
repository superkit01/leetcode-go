package hard

import "sort"

func orderlyQueue(s string, k int) string {

	if k > 1 {
		temp := []byte(s)
		sort.Slice(temp,
			func(i, j int) bool {
				return temp[i] < temp[j]
			})
		return string(temp)
	} else {
		min := s
		temp := s
		for i := 0; i < len(s); i++ {
			temp = temp[1:] + temp[:1]
			if temp < min {
				min = temp
			}
		}

		return min
	}

}
