package middle

import (
	"math"
	"sort"
)

// 排序+区间合并
func CountWays(ranges [][]int) int {
	if len(ranges) == 1 {
		return 2
	}

	sort.Slice(ranges, func(x, y int) bool {
		return ranges[x][0]-ranges[y][0] < 0
	})

	conbineRanges := make([][]int, 0)

	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		if current[1] < ranges[i][0] {
			conbineRanges = append(conbineRanges, current)
			current = ranges[i]
			continue
		} else {
			current = []int{current[0], int(math.Max(float64(current[1]), float64(ranges[i][1])))}
		}
	}

	conbineRanges = append(conbineRanges, current)

	return quickMultiply(2, len(conbineRanges))

}

// x ** n  快速幂
func quickMultiply(x, n int) int {
	//n =  0b101010010
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return quickMultiply(x, n-1) * x % 1000000007
	} else {
		return quickMultiply(x, n>>1) * quickMultiply(x, n>>1) % 1000000007
	}

}
