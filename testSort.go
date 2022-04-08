package main

import (
	"fmt"
	"leecode-go/sort"
)

func main() {
	nums := []int{5, 7, 3, 6, 2}
	sort.QuickSort(nums, 0, 4)
	fmt.Printf("%v \n", nums)
}
