package week398

import (
	"strconv"
)

func SumDigitDifferences(nums []int) int64 {
	n := len(strconv.Itoa(nums[0]))
	// i:个位数 、十位数、百位数、...
	// j:0,1,2,3...
	cnt := make([][]int, n)
	for i := range cnt {
		cnt[i] = make([]int, 10)
	}
	ans := int64(0)
	for i, v := range nums {
		temp := v
		pos := 0
		for temp > 0 {
			letter := temp % 10
			ans += int64(i - cnt[pos][letter])

			cnt[pos][letter]++
			temp = temp / 10
			pos++
		}
	}

	return int64(ans)

}
