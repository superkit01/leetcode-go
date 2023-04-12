package middle

import (
	"sort"
)

func NumberOfWeakCharacters(properties [][]int) int {

	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] > properties[j][0] {
			return true
		} else if properties[i][0] < properties[j][0] {
			return false

		} else {
			if properties[i][1] < properties[j][1] {
				return true
			} else {
				return false
			}
		}
	})
	result := 0
	max := properties[0][1]
	for i := 1; i < len(properties); i++ {
		if properties[i][1] < max {
			result++
		}
		if properties[i][1] > max {
			max = properties[i][1]
		}
	}

	return result

}
