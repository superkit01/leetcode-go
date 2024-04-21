package lcr

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals)==1 {
		return intervals
	}

	sort.Slice(intervals, func(i,j int)bool {
		if intervals[i][0]< intervals[j][0]{
			return true
		}else{
			return false
		}
	})


	ans:=make([][]int,0)

	current:=intervals[0]


	for i:=1;i<len(intervals);i++ {
		if intervals[i][0] > current[1] {
			ans=append(ans, current)
			current=intervals[i]
			continue
		}

		current[0]=int(math.Min(float64(current[0]),float64(intervals[i][0])))
		current[1]=int(math.Max(float64(current[1]),float64(intervals[i][1])))
	}

	ans=append(ans, current)
	
	return ans
}