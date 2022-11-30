package easy

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	arr = arr[len(arr)/20 : len(arr)*19/20]
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	return float64(sum) / float64(len(arr))

}
