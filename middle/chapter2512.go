package middle

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	fb := make(map[string]int, 0)

	for _, v := range positive_feedback {
		fb[v] = 3
	}

	for _, v := range negative_feedback {
		fb[v] = -1
	}

	reportValue := make([]int, 0)
	for _, v := range report {
		words := strings.Split(v, " ")
		value := 0
		for _, w := range words {
			if _, ok := fb[w]; ok {
				value += fb[w]
			}
		}
		reportValue = append(reportValue, value)
	}

	type Score struct {
		Value int
		Id    int
	}

	scoreList := make([]Score, 0)

	for i := 0; i < len(reportValue); i++ {
		scoreList = append(scoreList, Score{reportValue[i], student_id[i]})
	}

	sort.Slice(scoreList, func(x, y int) bool {
		if scoreList[x].Value > scoreList[y].Value {
			return true
		} else if scoreList[x].Value < scoreList[y].Value {
			return false
		} else {
			return scoreList[x].Id < scoreList[y].Id
		}
	})

	topK := scoreList[0:k]

	result := make([]int, 0)

	for _, v := range topK {
		result = append(result, v.Id)
	}
	return result

}
