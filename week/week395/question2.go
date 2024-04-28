package week395

import (
	"math"
	"sort"
)

func MinimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	min := math.MaxInt32
	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			flag, value := judgeEqual(nums1, nums2, i, j)
			if flag && min > value {
				min = value
			}
		}
	}
	return min

}

func judgeEqual(nums1 []int, nums2 []int, i, j int) (bool, int) {
	m := 0
	n := 0
	flag := false
	value := 0
	for m < len(nums1) {
		if m == i || m == j {
			m++
			continue
		}
		if !flag {
			value = nums2[n] - nums1[m]
			m++
			n++
			flag=true
			continue
		}

		if nums2[n]-nums1[m] != value {
			return false, 0
		}
		m++
		n++

	}

	return true, value

}
