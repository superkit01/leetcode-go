package sort

func InsertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

outer:
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				continue outer
			}

		}
	}

}
