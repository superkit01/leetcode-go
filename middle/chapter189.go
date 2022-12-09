package middle

func rotate1(nums []int, k int) {
	k = k % len(nums)
	reverse2(nums, 0, len(nums)-1)
	reverse2(nums, 0, k-1)
	reverse2(nums, k, len(nums)-1)

}

func reverse2(nums []int, start int, end int) {
	left := start
	right := end
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

}
