package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	tmp := make([]int, m+n)
	i := 0
	j := 0
	k := 0
	for {
		if k==m+n{
			break
		}
		if i >= m {
			tmp[k] = nums2[j]
			j++
		} else if j >= n {
			tmp[k] = nums1[i]
			i++
		} else if nums1[i] <= nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}

	copy(nums1,tmp)

}

//func main()  {
//	merge([]int{1,2,3,0,0}, 3,  []int{2,5,6},3)
//}
