package easy

func findIndices(nums []int, indexDifference int, valueDifference int) []int {

	if len(nums) < indexDifference {
		return []int{-1, -1}
	}

	for i := 0; i < len(nums)-indexDifference; i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			if nums[i]-nums[j] >= valueDifference || nums[i]-nums[j] <= -valueDifference {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
