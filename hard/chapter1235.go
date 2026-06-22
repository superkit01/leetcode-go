package hard

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	if n == 1 {
		return profit[0]
	}

	//重组
	timeProfit := make([][]int, n)
	for i := 0; i < n; i++ {
		timeProfit[i] = []int{startTime[i], endTime[i], profit[i]}
	}

	//按照结束时间升序排序
	sort.Slice(timeProfit, func(i, j int) bool {
		return timeProfit[i][1] < timeProfit[j][1]
	})

	//01背包 是否采用第i个计划得到的最大利润
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		//选择当前序列的工作计划
		current := timeProfit[i-1]
		//二分查找
		index := binarySearch(timeProfit, current[0])

		dp[i] = max(dp[i-1], dp[index]+current[2])
	}

	return dp[n]
}

//二分查找
func binarySearch(timeProfit [][]int, target int) int {
	l := 0
	r := len(timeProfit)

	for l < r {
		mid := (l + r) / 2
		//心智模型 l左边都是小于等于指定值 ==> l为大于指定值的第一个位置
		if timeProfit[mid][1] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
