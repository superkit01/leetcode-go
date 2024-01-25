package middle

import (
	"math"
	"sort"
)

func maximumSwap(num int) int {
	arr := make([]int, 0)
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10
	}

	newArr := make([]int, len(arr))
	copy(newArr, arr)

	sort.Ints(newArr)
	for i := len(newArr) - 1; i >= 0; i-- {
		if newArr[i] != arr[i] {

		inner:
			for j := 0; j < i; j++ {
				if arr[j] == newArr[i] {
					arr[j] = arr[i]
					break inner
				}
			}
			arr[i] = newArr[i]
			break
		}
	}

	result := 0

	for i := 0; i < len(arr); i++ {
		result += arr[i] * int(math.Pow(float64(10), float64(i)))
	}
	return result

}
