package middle

//输入：nums = [4,5,6,7,0,1,2], target = 0

//func main() {
//	search([]int{3, 1}, 1)
//}

func search(nums []int, target int) int {
	tmp := rotate(nums)

	if tmp == -1 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	if nums[0] > target {
		return binarySearch(nums, tmp+1, len(nums)-1, target)

	} else {
		return binarySearch(nums, 0, tmp, target)
	}

}

func binarySearch(nums []int, min int, max int, target int) int {

	for {

		mid := (max + min) / 2
		if min > max {
			return -1
		}
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			return mid
		}

	}

}

func rotate(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return -1
	}

	min := 0
	max := len(nums) - 1

	for {
		mid := (max + min) / 2
		if max-min <= 1 {
			return min
		}
		if nums[min] <= nums[mid] {
			min = mid
		} else {
			max = mid
		}
	}
}
