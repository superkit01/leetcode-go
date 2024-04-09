package easy

func maximumCount(nums []int) int {
	l := 0
	r := len(nums) - 1

	//小于0的最大位置
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= 0 {
			r = mid - 1
			continue
		}

		if nums[mid] < 0 {
			l = mid + 1
		}
	}
	//r

	m := 0
	n := len(nums) - 1

	//大于0的最大位置
	for m <= n {
		mid := (m + n) / 2
		if nums[mid] <= 0 {
			m = mid + 1
		}

		if nums[mid] > 0 {
			n = mid - 1
		}
	}

	//m
	if len(nums)-m > r+1 {
		return len(nums) - m
	}else{
		return r+1
	}
}
