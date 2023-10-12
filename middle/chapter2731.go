package middle

import (
	"sort"
)

func sumDistance(nums []int, s string, d int) int {

	for i := 0; i < len(nums); i++ {
		if s[i] == 'L' {
			nums[i] = nums[i] - d
		} else {
			nums[i] = nums[i] + d
		}
	}
	var result int64 = 0

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		temp := (((int64(nums[i]) - int64(nums[i-1])) % 1000000007) * ((int64(len(nums) - i)) % 1000000007) * (int64(i) % 1000000007)) % 1000000007
		result += temp
	}

	return int(result % 1000000007)

}
