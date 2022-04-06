package main

func findMaxLength(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	temp := make(map[int]int)
	temp[0] = -1
	sum := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := temp[sum]; ok {
			if i-v > max {
				max = i - v
			}
		} else {
			temp[sum] = i
		}
	}

	return max
}

//func main() {
//	findMaxLength([]int{0, 1})
//}
