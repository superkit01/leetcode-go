package week399

import (
	"math"
)

func NumberOfPairs(nums1 []int, nums2 []int, k int) int64 {

	cnt := map[int]int{}
	for _, k1 := range nums1 {
		if k1%k != 0 {
			continue
		}
		k1 = k1 / k
		temp := int(math.Sqrt(float64(k1)))
		for i := 1; i <= temp; i++ {
			if k1%i == 0 {
				cnt[i]++
				cnt[k1/i]++
			}
		}
		if temp*temp == k1 {
			cnt[temp]--
		}
	}

	ans := int64(0)

	for _, v := range nums2 {
		ans += int64(cnt[v])

	}

	return ans

}
