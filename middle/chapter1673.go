package middle

func MostCompetitive(nums []int, k int) []int {
	ans := make([]int, 0)
outer:
	for i := 0; i < len(nums); i++ {
		for {
			if len(ans) == 0 {
				ans = append(ans, nums[i])
				continue outer
			}

			if len(ans)+len(nums)-i == k {
				ans = append(ans, nums[i])
				continue outer
			}

			if ans[len(ans)-1] <= nums[i] {
				ans = append(ans, nums[i])
				continue outer
			}

			ans = ans[0 : len(ans)-1]

		}
	}

	return ans[0:k]

}
