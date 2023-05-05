package easy

import "sort"

func hardestWorker(n int, logs [][]int) int {
	if len(logs) == 1 {
		return logs[0][0]
	}

	max := logs[0][1]
	user := make([]int, 0)
	user = append(user, logs[0][0])

	for i := 1; i < len(logs); i++ {
		temp := logs[i][1] - logs[i-1][1]
		if temp > max {
			max = temp
			user = make([]int, 0)
			user = append(user, logs[i][0])
		} else if temp == max {
			user = append(user, logs[i][0])
		}
	}

	sort.Ints(user)
	return user[0]

}
