package main

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	result = append(result, []int{})

	for i := 0; i < len(nums); i++ {

		for _, v := range result {
			temp := make([]int, 0)
			temp = append(temp, v...)
			temp = append(temp, nums[i])
			result = append(result, temp)
		}

	}

	return result

}
