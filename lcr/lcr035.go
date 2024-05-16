package lcr

import (
	"sort"
	"strconv"
	"strings"
)

func FindMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	ans := 1440
	for i := 0; i < len(timePoints)-1; i++ {
		temp := sub(timePoints[i], timePoints[i+1])
		if temp < ans {
			ans = temp
		}
	}

	last := 1440 + sub(timePoints[len(timePoints)-1], timePoints[0])
	if last < ans {
		ans = last
	}

	return ans
}

func sub(t1, t2 string) int {
	t2Arr := strings.Split(t2, ":")
	t1Arr := strings.Split(t1, ":")
	t20, _ := strconv.Atoi(t2Arr[0])
	t10, _ := strconv.Atoi(t1Arr[0])
	t21, _ := strconv.Atoi(t2Arr[1])
	t11, _ := strconv.Atoi(t1Arr[1])
	return (t20-t10)*60 + t21 - t11
}
