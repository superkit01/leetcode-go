package middle

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	//[[maxValue , maxIndex, minValue, minIndex]]
	maxMin := make([][4]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			maxMin[i] = [4]int{nums[i], i, nums[i], i}
			continue
		}

		if nums[i] > maxMin[i+1][0] {
			maxMin[i][0] = nums[i]
			maxMin[i][1] = i
		} else {
			maxMin[i][0] = maxMin[i+1][0]
			maxMin[i][1] = maxMin[i+1][1]
		}

		if nums[i] < maxMin[i+1][2] {
			maxMin[i][2] = nums[i]
			maxMin[i][3] = i
		} else {
			maxMin[i][2] = maxMin[i+1][2]
			maxMin[i][3] = maxMin[i+1][3]
		}
	}

	for i := 0; i < len(nums)-indexDifference; i++ {
		j := i + indexDifference

		if nums[i]-maxMin[j][2] >= valueDifference {

			return []int{i, maxMin[j][3]}
		}

		if maxMin[j][0]-nums[i] >= valueDifference {
			return []int{i, maxMin[j][1]}
		}
	}

	return []int{-1, -1}

}
