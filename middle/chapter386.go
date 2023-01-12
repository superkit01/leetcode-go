package middle

import (
	"fmt"
	"strconv"
)

func lexicalOrder(n int) []int {
	result := make([]int, 0)

	for i := 1; i <= 9; i++ {
		digui1(i, &result, n)

	}

	return result

}

func digui1(val int, result *[]int, n int) {
	if val > n {
		return
	}

	*result = append(*result, val)
	for i := 0; i <= 9; i++ {
		temp, _ := strconv.Atoi(fmt.Sprint(val) + fmt.Sprint(i))
		digui1(temp, result, n)
	}
}
