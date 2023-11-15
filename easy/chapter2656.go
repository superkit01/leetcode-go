package easy

func maximizeSum(nums []int, k int) int {
	max := 0
	for _, v := range nums {
		if max < v {
			max = v
		}
	}

	return max*k + (k-1)*k/2

}
