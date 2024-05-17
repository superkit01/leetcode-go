package middle

func getWinner(arr []int, k int) int {
	max := arr[0]
	count := 0
	for i := 1; i < len(arr); i++ {
		if count == k {
			break
		}
		if max > arr[i] {
			count++
		} else {
			max = arr[i]
			count = 1
		}
	}
	return max
}
