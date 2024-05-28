package dweek131

func duplicateNumbersXOR(nums []int) int {
	cnt := make([]int, 51)
	ans := 0
	for _, v := range nums {
		cnt[v]++
		if cnt[v] == 2 {
			ans = ans ^ v
		}
	}
	return ans
}
