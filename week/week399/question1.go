package week399

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	for i := 0; i < len(nums2); i++ {
		nums2[i] = nums2[i] * k
	}
	ans := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%nums2[j] == 0 {
				ans++
			}
		}
	}
	return ans

}
