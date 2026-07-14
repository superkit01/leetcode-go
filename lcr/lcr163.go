package lcr

import (
	"fmt"
	"math"
)

func FindKthNumber(k int) int {
	//1-9   9	1
	//10-99  90		2
	//100-999 900		3
	//1000-9999 9000		4
	//10000-99999 90000			5
	//100000-999999 900000			6
	//1000000-9999999 9000000			7
	//10000000-99999999 90000000			8
	//100000000-999999999 900000000				9
	//1000000000-2147483647 1147483648				10

	digitCnt := []int{0, 9, 90, 900, 9000, 90000, 900000, 9000000, 90000000, 900000000, math.MaxInt32 - 1000000000 + 1}

	i := 1
	for {
		if i*digitCnt[i] >= k {
			break
		}
		k -= i * digitCnt[i]
		i++
	}
	//i=2 k=2
	if k%i == 0 {
		v := int(math.Pow(10, float64(i-1))) + k/i - 1
		fmt.Printf("v: %v", v)
		return v % 10
	} else {
		v := int(math.Pow(10, float64(i-1))) + k/i
		index := k % i
		fmt.Printf("v: %v, index: %v", v, index)
		return (v / int(math.Pow(10, float64(i-index)))) % 10
	}

}
