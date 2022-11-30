package middle

import (
	"math"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[1] {
			result = append(result, intervals[i:]...)
			break
		} else if intervals[i][1] < newInterval[0] {
			result = append(result, intervals[i])
		} else {
			newInterval[0] = int(math.Min(float64(intervals[i][0]), float64(newInterval[0])))
			newInterval[1] = int(math.Max(float64(intervals[i][1]), float64(newInterval[1])))
		}

	}
	result = append(result, newInterval)

	sort.Slice(result, func(i, j int) bool {
		if result[i][0]-result[j][0] > 0 {
			return false
		} else {
			return true
		}
	})

	return result

}

//func main() {
//	intervals := [][]int{{1, 5}}
//	fmt.Printf("%v \n", insert(intervals, []int{6, 8}))
//}
