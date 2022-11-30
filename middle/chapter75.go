package middle

func sortColors(nums []int) {
	p0 := 0
	p2 := len(nums) - 1

	for i := 0; i <= p2; {
		if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		} else if nums[i] == 1 {
			i++
		}
	}

}

// func main() {
// 	sortColors([]int{2, 0, 1})
// }
