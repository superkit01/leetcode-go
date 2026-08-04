package easy

import "math"

func findMissingElements(nums []int) []int {
	mi, mx := math.MaxInt32, math.MinInt32
	bit := make([]int, 1001)

	for _, v := range nums {
		if v < mi {
			mi = v
		}
		if v > mx {
			mx = v
		}

		bit[v] = 1
	}

	ans := make([]int, mx-mi+1-len(nums))
	index := 0
	for i := mi; i <= mx; i++ {
		if bit[i] == 0 {
			ans[index] = i
			index++
		}
	}
	return ans

}
