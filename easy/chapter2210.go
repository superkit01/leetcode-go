package easy

func countHillValley(nums []int) int {
	temp := make([]int, 0)
	temp = append(temp, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else {
			temp = append(temp, nums[i])
		}
	}

	if len(temp) < 3 {
		return 0
	}
	count := 0

	for i := 1; i < len(temp)-1; i++ {
		if temp[i] > temp[i-1] && temp[i] > temp[i+1] {
			count += 1
		}
		if temp[i] < temp[i-1] && temp[i] < temp[i+1] {
			count += 1
		}
	}
	return count
}
