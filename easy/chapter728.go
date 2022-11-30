package easy

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
outer:
	for i := left; i <= right; i++ {
		temp := i
		for temp > 0 {
			if temp%10 == 0 {
				continue outer
			}
			if i%(temp%10) != 0 {
				continue outer
			}
			temp = temp / 10
		}
		result = append(result, i)

	}
	return result
}
