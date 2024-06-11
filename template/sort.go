package template

// 选择排序
// 每次选择 i+1 ~ len(nums) 最小值放入当前位置
func SelectSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		min := i //最小值对应的index
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

// 冒泡排序
// 每次 len(nums)~ i 依次交换较小值冒泡到i位置
func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

//插入  0 ~ i-1位已排序
//将i位置元素插入到合适的位置
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
