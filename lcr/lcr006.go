package lcr

//二分
func twoSum(numbers []int, target int) []int {
	for i, v := range numbers {

		l := i + 1
		r := len(numbers) - 1

		for l <= r {
			mid := (l + r) / 2
			if numbers[mid] == target-v {
				return []int{i, mid}
			}
			if numbers[mid] > target-v {
				r = mid - 1
			}
			if numbers[mid] < target-v {
				l = mid + 1
			}
		}

	}
	return []int{-1, -1}

}
