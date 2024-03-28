package sort
//快速
func QuickSort(nums []int, L int, R int) {
	if len(nums) <= 1 {
		return
	}
	if L >= R {
		return
	}
	left := L
	right := R

	pivot := nums[left]

	for left < right {
		for left < right && nums[right] > pivot {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
		}

		for left < right && nums[left] < pivot {
			left++
		}
		if left < right {
			nums[right] = nums[left]
			right--
		}

	}
	nums[left] = pivot

	QuickSort(nums, L, left-1)
	QuickSort(nums, left+1, R)

}
