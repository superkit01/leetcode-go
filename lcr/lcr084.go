package lcr

import "sort"

func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)

	ans := [][]int{}

	var dfs func(num, current []int)

	dfs = func(num, current []int) {
		if len(num) == 0 {
			ans = append(ans, current)
			return
		}

		for i := 0; i < len(num); i++ {
			if i > 0 && num[i] == num[i-1] {
				continue
			}

			temp := append(append([]int{}, num[0:i]...), num[i+1:]...)
			currTemp := append(append([]int{}, current...), num[i])
			dfs(temp, currTemp)
		}
	}

	dfs(nums, []int{})
	return ans

}
