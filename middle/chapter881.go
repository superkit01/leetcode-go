package middle

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i := 0
	j := len(people) - 1

	ans := 0
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
			j--
		} else {
			j--
		}
		ans++
	}
	return ans
}
