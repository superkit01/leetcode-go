package hard

import "sort"

func findMinimumTime(tasks [][]int) int {

	//按照结束时间从小到大排序
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i][1] == tasks[j][1] {
			return tasks[i][0] < tasks[j][0]
		}
		return tasks[i][1] < tasks[j][1]
	})

	//遍历  优先在结束时间点运行
	work := [2001]int{}
	for i, v := range tasks {
		duration := tasks[i][2]
		for j := v[0]; j <= v[1]; j++ {
			if work[j] == 1 {
				duration--
			}
		}

		for j := v[1]; j >= v[0]; j-- {
			if duration <= 0 {
				break
			}
			if work[j] == 0 {
				work[j] = 1
				duration--
			}
		}
	}

	ans := 0
	for _, v := range work {
		if v == 1 {
			ans++
		}
	}
	return ans
}
