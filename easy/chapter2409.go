package easy

import (
	"strconv"
	"strings"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	if arriveAlice > leaveBob || leaveAlice < arriveBob {
		return 0
	}

	cache := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	start := ""
	if arriveAlice > arriveBob {
		start = arriveAlice
	} else {
		start = arriveBob
	}

	end := ""
	if leaveAlice < leaveBob {
		end = leaveAlice
	} else {
		end = leaveBob
	}

	startSlice := strings.Split(start, "-")
	endSlice := strings.Split(end, "-")
	if startSlice[0] == endSlice[0] {
		start1, _ := strconv.Atoi(startSlice[1])
		end1, _ := strconv.Atoi(endSlice[1])
		return end1 - start1 + 1
	} else {
		start0, _ := strconv.Atoi(startSlice[0])
		start1, _ := strconv.Atoi(startSlice[1])
		end0, _ := strconv.Atoi(endSlice[0])
		end1, _ := strconv.Atoi(endSlice[1])
		days := cache[start0-1] - start1 + 1 + end1
		if end0-start0 > 1 {
			for i := start0 + 1; i < end0; i++ {
				days += cache[i-1]
			}

		}
		return days
	}

}
