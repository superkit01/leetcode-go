package hard

import (
	"math"
	"sort"
)

// [[3, 2], [5, 1], [10, 1]], 2
// TLE  记忆化搜索
func FindMaximumElegance(items [][]int, k int) int64 {
	category := map[int]int{}

	var dfs func(index int, remain int, profit int64) (int64, int)

	dfs = func(index int, remain int, profit int64) (int64, int) {
		if remain == 0 {
			return profit, len(category)
		}

		if index >= len(items) {
			return math.MinInt64, 0
		}

		//选当前位置
		category[items[index][1]]++
		p1, c1 := dfs(index+1, remain-1, profit+int64(items[index][0]))

		if category[items[index][1]] == 1 {
			delete(category, items[index][1])
		}
		//不选当前位置
		p2, c2 := dfs(index+1, remain, profit)
		if p1+int64(c1*c1) > p2+int64(c2*c2) {
			return p1, c1
		} else {
			return p2, c2
		}
	}
	p, c := dfs(0, k, int64(0))
	return p + int64(c*c)
}

// 反悔贪心
func FindMaximumEleganceII(items [][]int, k int) int64 {
	//从大到小排序
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})

	stack := []int{}
	category := map[int]int{}
	profit := int64(0)
	for i := 0; i < k; i++ {
		if category[items[i][1]] > 0 {
			stack = append(stack, items[i][0]) //item本身排序好的，所以这边也是从大到小排序的
		}
		category[items[i][1]]++
		profit += int64(items[i][0])
	}
	ans := profit + int64(len(category)*len(category))

	for i := k; i < len(items); i++ {
		if _, ok := category[items[i][1]]; ok {
			continue
		}
		if len(stack) == 0 {
			continue
		}

		//替换stack中最小的
		category[items[i][1]]++
		profit += int64(items[i][0])
		profit -= int64(stack[len(stack)-1])
		stack = stack[0 : len(stack)-1]

		ans = maxInt64(ans, profit+int64(len(category)*len(category)))
	}

	return ans
}

func maxInt64(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}
