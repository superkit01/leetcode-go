package lcr

func Permute(nums []int) [][]int {
	ans := [][]int{}

	var dfs func(num []int, current []int)
	dfs = func(num []int, current []int) {
		if len(num) == 0 {
			ans = append(ans, current)
			return
		}

		for i, v := range num {
			temp := append(append([]int{}, num[0:i]...), num[i+1:]...)
			currTemp := append([]int{}, current...)
			currTemp = append(currTemp, v)
			dfs(temp, currTemp)
		}
	}

	dfs(nums, []int{})
	return ans
}
