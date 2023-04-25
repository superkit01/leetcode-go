package easy

import "sort"

type P struct {
	name   string
	height int
}

func SortPeople(names []string, heights []int) []string {
	people := make([]P, 0)
	for i, k := range names {
		people = append(people, P{name: k, height: heights[i]})
	}

	sort.Slice(people, func(i, j int) bool { return people[i].height > people[j].height })

	result := make([]string, 0)

	for _, v := range people {
		result = append(result, v.name)
	}
	return result

}
