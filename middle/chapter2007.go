package middle

import "sort"

func FindOriginalArray(changed []int) []int {
	if len(changed) == 0 || len(changed)%2 == 1 {
		return []int{}
	}
	odd := make([]int, 0)
	even := make([]int, 0)
	evenMap := make(map[int][]int, 0)

	for i, v := range changed {
		if v%2 == 1 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
			evenMap[v] = append(evenMap[v], i)
		}
	}

	sort.Ints(even)

	for _, v := range odd {
		if _, ok := evenMap[2*v]; !ok || len(evenMap[2*v]) <= 0 {
			return []int{}
		}
		evenMap[2*v] = evenMap[2*v][1:]
	}

	for _, v := range even {
		if len(evenMap[v]) == 0 {
			continue
		}
		evenMap[v] = evenMap[v][1:]

		if _, ok := evenMap[2*v]; !ok || len(evenMap[2*v]) <= 0 {
			return []int{}
		}
		odd = append(odd, v)
		evenMap[2*v] = evenMap[2*v][1:]
	}

	return odd

}
