package middle

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	result := make([]int, 0)

	for _, v := range spells {
		index := bs(potions, v, success)
		result = append(result, len(potions)-index)
	}
	return result

}

func bs(potions []int, v int, success int64) int {

	if int64(potions[0])*int64(v) >= success {
		return 0
	}

	if int64(potions[len(potions)-1])*int64(v) < success {
		return len(potions)
	}

	left := 0
	right := len(potions) - 1

	for left <= right {
		mid := (right + left) / 2

		if int64(potions[mid])*int64(v) < success {
			left = mid + 1
		} else if int64(potions[mid])*int64(v) >= success {
			right = mid - 1
		}

	}
	return left
}
